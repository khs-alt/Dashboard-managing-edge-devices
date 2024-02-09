package command

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/crypto/ssh"
)

// ExecuteCommand executes a local command and returns its output.
func ExecuteCommand(command string, args []string) (string, error) {
	// Create a Command struct using exec.Command function.
	cmd := exec.Command(command, args...)

	// Create a buffer to capture the standard output of the command.
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// Execute the command.
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Return the output of the command as a string.
	return stdout.String(), nil
}

// ExecuteSSHCommand uses the given SSH details to execute a command on a remote system and returns the result.
func ExecuteSSHCommand(host, port, user, privateKeyPath, command string) (string, error) {
	// Authenticate using a private key.
	key, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return "", fmt.Errorf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Use a more secure method in actual use.
	}

	// Create an SSH connection.
	connection, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return "", fmt.Errorf("failed to dial: %v", err)
	}
	defer connection.Close()

	// Create a session.
	session, err := connection.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// Execute the command.
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	err = session.Run(command)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %v", err)
	}

	return stdoutBuf.String(), nil
}
