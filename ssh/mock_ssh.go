// Copyright 2015 Apcera Inc. All rights reserved.

package ssh

import (
	"io"
	"time"
)

// Connect calls the mocked connect.
func (c *MockSSHClient) Connect() error {
	if c.MockConnect != nil {
		return c.MockConnect()
	}
	return ErrNotImplemented
}

// Disconnect calls the mocked disconnect.
func (c *MockSSHClient) Disconnect() {
	if c.MockDisconnect != nil {
		c.MockDisconnect()
		return
	}
	return
}

// Download calls the mocked download.
func (c *MockSSHClient) Download(src io.WriteCloser, dst string) error {
	if c.MockDownload != nil {
		return c.MockDownload(src, dst)
	}
	return ErrNotImplemented
}

// Run calls the mocked run
func (c *MockSSHClient) Run(command string, stdout io.Writer, stderr io.Writer) error {
	if c.MockRun != nil {
		return c.MockRun(command, stdout, stderr)
	}
	return ErrNotImplemented
}

// Upload calls the mocked upload
func (c *MockSSHClient) Upload(src io.Reader, dst string, mode uint32) error {
	if c.MockUpload != nil {
		return c.MockUpload(src, dst, mode)
	}
	return ErrNotImplemented
}

// Validate calls the mocked validate.
func (c *MockSSHClient) Validate() error {
	if c.MockValidate != nil {
		return c.MockValidate()
	}
	return ErrNotImplemented
}

// WaitForSSH calls the mocked WaitForSSH
func (c *MockSSHClient) WaitForSSH(maxWait time.Duration) error {
	if c.MockWaitForSSH != nil {
		return c.MockWaitForSSH(maxWait)
	}
	return ErrNotImplemented
}

// SetSSHPrivateKey calls the mocked SetSSHPrivateKey
func (c *MockSSHClient) SetSSHPrivateKey(s string) {
	if c.MockSetSSHPrivateKey != nil {
		c.MockSetSSHPrivateKey(s)
		return
	}
	return
}

// SetSSHPassword calls the mocked SetSSHPassword
func (c *MockSSHClient) SetSSHPassword(s string) {
	if c.MockSetSSHPassword != nil {
		c.MockSetSSHPassword(s)
		return
	}
	return
}
