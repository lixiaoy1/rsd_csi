package rsd_csi

import (
    "fmt"
    "os/exec"
    "strings"
    "github.com/golang/glog"
)

func execCommand(command string, args []string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	return cmd.CombinedOutput()
}

func connectRSDVolume(initiator string, target string, target_ip string) (error) {
    cmd := "nvme"
    args := []string{"connect", "-t", "rdma", "-n", target, "-a", target_ip, "-s", "4420", "-q", initiator}
    glog.Infof("Nvme command args: %v", args)
    err := exec.Command(cmd, args...).Run()
    return err
}

func getDevicePath(target string) (string){
    cmd := exec.Command("nvme", "list")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return ""
    }

    nvme_list := string(out)
    fmt.Println("nvme_list:\n%s", nvme_list)
    nvmes := strings.Split(nvme_list, "\n")

    for i, n := range nvmes {
        if i > 1 {
            devicePath := strings.Split(n, " ")[0]
            fmt.Printf("%q\n", devicePath)
            cmd = exec.Command("nvme", "id-ctrl", devicePath)
            out, err = cmd.CombinedOutput()
            if err != nil {
                continue
            }
            id_list:= string(out)
            ids := strings.Split(id_list, "\n")
            for _, q := range ids {
                //string compare with nqn
                if strings.Split(strings.Split(q, ": ")[0], " ")[0] != "subnqn" {
                    continue
                }
                nqn := strings.Split(q, ": ")[1]
                fmt.Printf("NQN: %q\n", nqn)
                if nqn == target {
                    fmt.Printf("Returned devicePath: %q\n", devicePath)
                    return devicePath
                }
            }
        }
    }

    return ""
}

