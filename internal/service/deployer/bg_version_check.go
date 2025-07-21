package deployer

import (
	"fmt"
	"noize_metter/internal/utils"
	"os"
	"os/exec"
	"time"
)

func (s *Service) fetchLatestVersionLoop() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			s.fetchLatestVersion()
		}
	}
}

// fetchLatestVersion compare current app version with the latest version from the repository.
// If they differ, it logs, pull changes, and triggers a self-deploy.
func (s *Service) fetchLatestVersion() {
	lastVersion, err := s.repo.FetchLatestVersion(s.ctx)
	if err != nil {
		s.log.Error("failed to fetch latest version", err)
		return
	}
	appVersion := utils.GetLastCommitHash()
	if lastVersion == appVersion {
		return
	}

	cmd := exec.CommandContext(s.ctx, "bash", "-c", fmt.Sprintf("cd %s && make self_deploy", s.cfg.AppFolder))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		s.log.Error("new version self_deploy failed", err)
		_ = s.notify.SendInfoMessage("self_deploy failed", err.Error())
		return
	}
	_ = s.notify.SendInfoMessage("app restart triggered", fmt.Sprintf("current: %s, new: %s", appVersion, lastVersion))
	s.rst <- os.Interrupt
}
