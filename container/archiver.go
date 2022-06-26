package container

import (
	"errors"
	"os/exec"
)

func Compress(srcPath string, dstFile string) error {
	dstFile = dstFile + ".tar"
	_, err := exec.Command("tar", "-czf", dstFile, "-C", srcPath, ".").CombinedOutput()
	return err
}

func Decompress(srcFile string, dstPath string) error {
	srcFile = srcFile + ".tar"
	if !ExistFile(srcFile) {
		return errors.New("file not exist: " + srcFile)
	}
	_, err := exec.Command("tar", "-xvf", srcFile, "-C", dstPath).CombinedOutput()
	return err
}
