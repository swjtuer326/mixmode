package sophon

import (
	"fmt"
	"os/exec"

	"sophgo.com/mixmode/utils/common"
)

func TestCDMA() (string, error) {
	out, err := exec.Command("test_cdma_perf").Output()
	if err != nil {
		common.Log.Warn(err)
		return "", err
	}
	fmt.Println(string(out))
	return string(out), nil
	// return `
	// bbdbd
	// wjncjn`, nil
}

func TestGDMA() (string, error) {
	out, err := exec.Command("test_gdma_perf").Output()
	if err != nil {
		common.Log.Warn(err)
		return "", err
	}
	fmt.Println(string(out))
	return string(out), nil
}
