package notificator

import (
	"context"
	"fmt"
	"net/http"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
	"noize_metter/internal/utils"
	"noize_metter/internal/utils/requests"
	"strings"
	"sync"
	"time"
)

type SlackService struct {
	boxPreBuild string
	log         logger.AppLogger
	conf        *config.AppConfig
}

type slackBlock struct {
	Type     string    `json:"type"`
	BlockID  string    `json:"block_id,omitempty"`
	Text     *item     `json:"text,omitempty"`
	Fields   []item    `json:"fields,omitempty"`
	Elements []element `json:"elements,omitempty"`
}

type element struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	Elements []item `json:"elements,omitempty"`
}

type item struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

var (
	divider = slackBlock{Type: "divider"}
)

func NewSlackService(l logger.AppLogger, conf *config.AppConfig) *SlackService {
	return &SlackService{
		log:  l,
		conf: conf,
	}
}

func (s *SlackService) SendTaskErrMessage(service string, startedAt, finishedAt time.Time, errs ...error) error {
	hooksURLs := s.conf.SlackHookURLs
	if !shouldSend(hooksURLs) {
		for i := range errs {
			if errs[i] == nil {
				continue
			}
			s.log.Error("service job failed", errs[i])
		}
		return nil
	}
	errMsg := strings.Builder{}
	errBlock := slackBlock{
		Type:     "rich_text",
		BlockID:  "block1",
		Elements: make([]element, 0, len(errs)),
	}
	prefix := false
	if len(errs) > 1 {
		prefix = true
	}
	for i := range errs {
		if errs[i] == nil {
			continue
		}
		prefixText := ""
		if prefix {
			prefixText = fmt.Sprintf("error %d: ", i+1)
		}
		errBlock.Elements = append(errBlock.Elements, element{
			Type: "rich_text_section",
			Elements: []item{
				{
					Type: "text",
					Text: fmt.Sprintf("%s%s\n", prefixText, errs[i].Error()),
				},
			},
		})
		errMsg.WriteString(fmt.Sprintf("error %d: %s\n", i+1, errs[i].Error()))
	}
	if errMsg.Len() == 0 {
		return nil
	}

	return s.sendSlackMessage(hooksURLs, map[string][]slackBlock{
		"blocks": {
			getHeader(fmt.Sprintf(":bangbang: Service job failed: %s", service)),
			divider,
			errBlock,
			{
				Type: "section",

				Fields: []item{
					{
						Type: "mrkdwn",
						Text: "*started at*",
					},
					{
						Type: "mrkdwn",
						Text: "*finished at*",
					},
					{
						Type: "mrkdwn",
						Text: startedAt.Format(time.DateTime),
					},
					{
						Type: "mrkdwn",
						Text: fmt.Sprintf("%s (%s)", finishedAt.Format(time.DateTime), finishedAt.Sub(startedAt).String()),
					},
				},
			},
			divider,
			s.getContext(),
		},
	})
}

func (s *SlackService) SendInfoMessage(message string, args ...string) error {
	return s.sendInfoMessage(message, args...)
}

func (s *SlackService) sendInfoMessage(message string, args ...string) error {
	hooksURLs := s.conf.SlackHookURLs
	if !shouldSend(hooksURLs) {
		return nil
	}
	blocks := make([]slackBlock, 0, len(args)+3)
	blocks = append(blocks, getHeader(fmt.Sprintf(":warning: %s", message)))
	blocks = append(blocks, divider)
	if len(args) > 0 {
		for i := range args {
			blocks = append(blocks, slackBlock{
				Type: "section",
				Text: &item{
					Type: "mrkdwn",
					Text: args[i],
				},
			})
		}
		blocks = append(blocks, divider)
	}
	blocks = append(blocks, s.getContext())
	return s.sendSlackMessage(hooksURLs, map[string][]slackBlock{"blocks": blocks})
}

func (s *SlackService) sendSlackMessage(hookURLs []string, payload any) error {
	var (
		wg      sync.WaitGroup
		rwSlice = utils.NewRWSlice[error]()
	)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for i := range hookURLs {
		wg.Add(1)
		go func(hookURL string) {
			defer wg.Done()
			_, code, err := requests.PostCurl[any](ctx, hookURL, payload, map[string]string{"Content-Type": "application/json"})
			if code == http.StatusOK {
				return
			}
			if err != nil {
				rwSlice.Add(fmt.Errorf("failed to send notification to Slack: %w", err))
			}
			if code >= 300 {
				rwSlice.Add(fmt.Errorf("received non-2xx response from Slack: %d", code))
			}
		}(hookURLs[i])
	}
	wg.Wait()
	for _, err := range rwSlice.LoadAll() {
		return err
	}
	return nil
}

func getHeader(message string) slackBlock {
	return slackBlock{
		Type: "header",
		Text: &item{
			Type: "plain_text",
			Text: message,
		},
	}
}

func (s *SlackService) getContext() slackBlock {
	return slackBlock{
		Type: "context",
		Elements: []element{
			{
				Type: "mrkdwn",
				Text: fmt.Sprintf("box name: *%s*, box ip: *%s*, hash: *%s*", s.conf.BoxName, s.conf.BoxIP, utils.GetLastCommitHash()),
			},
		},
	}
}

func shouldSend(hooksURLs []string) bool {
	onlyEmpty := len(hooksURLs) == 0 || len(utils.ExcludeFromSlice(hooksURLs, []string{""})) == 0
	return !onlyEmpty
}
