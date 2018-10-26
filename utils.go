package rsd_csi

import (
    "fmt"
    "os/exec"
    "strings"
    "strconv"
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
    glog.Infof("getDevicePath from target: %s.", target)
    if target == "" {
       return ""
    }
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


func getDevicePath2(target string) (string){
    glog.Infof("getDevicePath2 from target: %s", target)
    cmd := exec.Command("nvme", "list")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return ""
    }

    nvme_max := 0
    nvme_list := string(out)
    nvmes := strings.Split(nvme_list, "\n")
    for i, n := range nvmes {
        if i > 1 {
            devicePath := strings.Split(n, " ")[0]
            //fmt.Printf("%q\n", devicePath)
            if devicePath == "" {
                continue
            }
            nvme_no, err := strconv.Atoi(strings.Split(strings.Split(devicePath, "nvme")[1], "n")[0])
            //fmt.Printf("__tingjie nvme_no: %d, nvme_max: %d", nvme_no, nvme_max)
            if (err == nil) && (nvme_no > nvme_max) {
                nvme_max = nvme_no
            }
        }
    }

    devicePathPlus := fmt.Sprintf("/dev/nvme%sn1", strconv.Itoa(nvme_max))
    fmt.Printf("Returned devicePathPlus: %q\n", devicePathPlus)
    return devicePathPlus
}

