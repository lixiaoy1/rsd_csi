package rsd_csi

import (
    "os/exec"
)

func execCommand(command string, args []string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	return cmd.CombinedOutput()
}

func connectRSDVolume(initiator string, target string, target_ip string) (string, error) {
    /*
    err := execCommand('usr/local/sbin/nvme', 'connect',
                       '-t', 'rdma',
                       '-n', target,
                       '-a', target_ip,
                       '-s', '4220',
                   '    -q', initiator)
                */
    return "path", nil
}
