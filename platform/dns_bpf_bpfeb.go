// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package platform

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadDns_bpf returns the embedded CollectionSpec for dns_bpf.
func loadDns_bpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Dns_bpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load dns_bpf: %w", err)
	}

	return spec, err
}

// loadDns_bpfObjects loads dns_bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*dns_bpfObjects
//	*dns_bpfPrograms
//	*dns_bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadDns_bpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadDns_bpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// dns_bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dns_bpfSpecs struct {
	dns_bpfProgramSpecs
	dns_bpfMapSpecs
}

// dns_bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dns_bpfProgramSpecs struct {
	TraceTcpSendmsg *ebpf.ProgramSpec `ebpf:"trace_tcp_sendmsg"`
	TraceUdpSendmsg *ebpf.ProgramSpec `ebpf:"trace_udp_sendmsg"`
}

// dns_bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type dns_bpfMapSpecs struct {
	Events *ebpf.MapSpec `ebpf:"events"`
}

// dns_bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadDns_bpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type dns_bpfObjects struct {
	dns_bpfPrograms
	dns_bpfMaps
}

func (o *dns_bpfObjects) Close() error {
	return _Dns_bpfClose(
		&o.dns_bpfPrograms,
		&o.dns_bpfMaps,
	)
}

// dns_bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadDns_bpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type dns_bpfMaps struct {
	Events *ebpf.Map `ebpf:"events"`
}

func (m *dns_bpfMaps) Close() error {
	return _Dns_bpfClose(
		m.Events,
	)
}

// dns_bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadDns_bpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type dns_bpfPrograms struct {
	TraceTcpSendmsg *ebpf.Program `ebpf:"trace_tcp_sendmsg"`
	TraceUdpSendmsg *ebpf.Program `ebpf:"trace_udp_sendmsg"`
}

func (p *dns_bpfPrograms) Close() error {
	return _Dns_bpfClose(
		p.TraceTcpSendmsg,
		p.TraceUdpSendmsg,
	)
}

func _Dns_bpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed dns_bpf_bpfeb.o
var _Dns_bpfBytes []byte
