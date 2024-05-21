package operatorsinfo

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestStartAnvilChain(t *testing.T) {
	pid, err := startAnvilChain()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Anvil chain started with PID", pid)

	curr_path, err := os.Getwd()
	changeDirectory(curr_path + "integration_test_deployment/lib/eigenlayer-middleware-offchain/subgraphs/BLSApkRegistry")

	_, err = startGraph()
	if err != nil {
		t.Fatal(err)
	}
	_, err = startSubgraph()
	if err != nil {
		t.Fatal(err)
	}

	// err = stopAnvilChain(pid)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

func startAnvilChain() (int, error) {
	return execCmd("anvil", []string{"--host", "0.0.0.0"}, []string{})
}

func startGraph() (int, error) {
	return execCmd("docker", []string{"compose", "up"}, []string{})
}

func startSubgraph() (int, error) {
	_, err := execCmd("graph", []string{"codegen"}, []string{})
	if err != nil {
		return 0, err
	}
	_, err = execCmd("graph", []string{"build"}, []string{})
	if err != nil {
		return 0, err
	}
	_, err = execCmd("npm", []string{"run", "create-local"}, []string{})
	if err != nil {
		return 0, err
	}
	return execCmd("npm", []string{"run", "deploy-local"}, []string{})
}

func stopAnvilChain(pid int) error {
	_, err := execCmd("kill", []string{"-9", fmt.Sprintf("%d", pid)}, []string{})
	return err
}

func execCmd(name string, args []string, envVars []string) (int, error) {
	cmd := exec.Command(name, args...)
	if len(envVars) > 0 {
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, envVars...)
	}
	var out bytes.Buffer
	var stderr bytes.Buffer
	// TODO: When these are uncommented, the deployer sometimes fails to start anvil
	// cmd.Stdout = &out
	// cmd.Stderr = &stderr

	fmt.Print("Running command: ", cmd.String())
	err := cmd.Run()
	fmt.Print("Running command: ", cmd.Process.Pid)
	if err != nil {
		return 0, fmt.Errorf("%s: %s", err.Error(), stderr.String())
	}
	fmt.Print(out.String())
	pid := cmd.Process.Pid
	return pid, nil
}

func changeDirectory(path string) {
	err := os.Chdir(path)
	if err != nil {
		log.Panicf("Failed to change directories. Error: %s", err)
	}

	newDir, err := os.Getwd()
	if err != nil {
		log.Panicf("Failed to get working directory. Error: %s", err)
	}
	log.Printf("Current Working Directory: %s\n", newDir)
}
