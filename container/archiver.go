package container

import (
	"errors"
	"os/exec"
)

func Compress(srcPath string, dstFile string) error {
	_, err := exec.Command("tar", "-czf", dstFile+".tar", "-C", srcPath, ".").CombinedOutput()
	return err
}

func Decompress(srcFile string, dstPath string) error {
	if !ExistFile(srcFile) {
		return errors.New("file not exist: " + srcFile)
	}
	_, err := exec.Command("tar", "-xvf", srcFile, "-C", dstPath).CombinedOutput()
	return err
}
