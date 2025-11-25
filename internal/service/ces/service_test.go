package ces_test

import (
	"context"
	"fmt"
	"noize_metter/internal/entities"
	testhelpers "noize_metter/internal/test_helpers"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewService(t *testing.T) {
	container := testhelpers.GetClean(t)

	// 1. create user
	// 2. get token
	// 3. get all mega boxes

	// http://10.20.10.101/grafana/login - not auth even after user created
	// http://10.20.10.101/swagger
	urlList := []string{
		// bellow works
		"http://10.20.10.101",
		"http://10.20.10.107",

		"http://10.20.10.116", // no errors
		"http://10.20.10.117", // no errors

		// ---------------------
		//"https://10.20.10.104",
		//"https://10.20.10.107",
		//// IMDC-2"
		//"https://10.20.10.116",
		//"https://10.20.10.117",

		// "http://10.20.10.104", // can create user, can autorize, endpoint serve 500x error
		// "http://10.20.10.107", // can't create user, got 400 for autorization
		//
		//"http://10.20.10.116", // can't create user, got 400 for autorization
		//"http://10.20.10.117", // can't create user, got 400 for autorization

		// IMDC1:
		//10.20.10.104 - https://taal-yfca1-hmi1.taal.intelliflexids.ca/device/WebRH
		//10.20.10.105 - https://taal-yfca1-hmi2.taal.intelliflexids.ca/device/WebRH
		//10.20.10.108 - https://taal-yfca1-hmi3.taal.intelliflexids.ca/device/WebRH
		//10.20.10.109 - https://taal-yfca1-hmi4.taal.intelliflexids.ca/device/WebRH
		//
		//IMDC2:
		//10.20.10.114 - https://taal-yfca2-hmi1.taal.intelliflexids.ca/device/WebRH
		//10.20.10.115 - https://taal-yfca2-hmi2.taal.intelliflexids.ca/device/WebRH
		//10.20.10.118 - https://taal-yfca2-hmi3.taal.intelliflexids.ca/device/WebRH
		//10.20.10.119 - https://taal-yfca2-hmi4.taal.intelliflexids.ca/device/WebRH
	}
	result := make([][]entities.MegaBox, 0, len(urlList))
	for _, baseURL := range urlList {
		println(fmt.Sprintf("--- testing %s", baseURL))
		boxes, err := container.ServiceCes.GetAllMegaBoxes(context.Background())
		if err != nil {
			t.Logf("failed to get all mega boxes from %s: %v", baseURL, err)
			continue
		}
		for _, b := range boxes {
			if b.Error != "" {
				t.Logf("mega box %d error: %s", b.MegaBoxID, b.Error)
			}
		}
		tanks, err := container.ServiceCes.GetAllTanks(context.Background())
		if err != nil {
			t.Logf("failed to get all tanks from %s: %v", baseURL, err)
			continue
		}
		for _, tank := range tanks {
			if tank.Error != "" {
				t.Logf("tank %d error: %s", tank.TankID, tank.Error)
			}
			fanInfo, err := container.ServiceCes.GetVFDFan(context.Background(), tank.TankID, "")
			if err != nil {
				t.Logf("failed to get fan info for tank %d: %v", tank.TankID, err)
			}
			if fanInfo != nil && fanInfo.Error != "" {
				t.Logf("fan info for tank %d error: %s", tank.TankID, fanInfo.Error)
			}
			pumpInfo, err := container.ServiceCes.GetVFDPump(context.Background(), tank.TankID, "")
			if err != nil {
				t.Logf("failed to get pump info for tank %d: %v", tank.TankID, err)
			}
			if pumpInfo != nil && pumpInfo.Error != "" {
				t.Logf("pump info for tank %d error: %s", tank.TankID, pumpInfo.Error)
			}
		}
		result = append(result, boxes)
	}
	require.Len(t, result, len(urlList))
}
