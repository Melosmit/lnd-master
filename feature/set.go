package feature

import (
	"math"

	"github.com/lightningnetwork/lnd/lnwire"
)

// Set is an enum identifying various feature sets, which separates the single
// feature namespace into distinct categories depending what context a feature
// vector is being used.
type Set uint8

const (
	// SetInit identifies features that should be sent in an Init message to
	// a remote peer.
	SetInit Set = iota

	// SetLegacyGlobal identifies features that should be set in the legacy
	// GlobalFeatures field of an Init message, which maintains backwards
	// compatibility with nodes that haven't implemented flat features.
	SetLegacyGlobal

	// SetNodeAnn identifies features that should be advertised on node
	// announcements.
	SetNodeAnn

	// SetInvoice identifies features that should be advertised on invoices
	// generated by the daemon.
	SetInvoice

	// SetInvoiceAmp identifies the features that should be advertised on
	// AMP invoices generated by the daemon.
	SetInvoiceAmp

	// setSentinel is used to mark the end of our known sets. This enum
	// member must *always* be the last item in the iota list to ensure
	// that validation works as expected.
	setSentinel
)

// valid returns a boolean indicating whether a set value is one of our
// predefined feature sets.
func (s Set) valid() bool {
	return s < setSentinel
}

// String returns a human-readable description of a Set.
func (s Set) String() string {
	switch s {
	case SetInit:
		return "SetInit"
	case SetLegacyGlobal:
		return "SetLegacyGlobal"
	case SetNodeAnn:
		return "SetNodeAnn"
	case SetInvoice:
		return "SetInvoice"
	case SetInvoiceAmp:
		return "SetInvoiceAmp"
	default:
		return "SetUnknown"
	}
}

// Maximum returns the maximum allowable value for a feature bit in the context
// of a set. The maximum feature value we can express differs by set context
// because the amount of space available varies between protocol messages. In
// practice this should never be a problem (reasonably one would never hit
// these high ranges), but we enforce these maximums for the sake of sane
// validation.
func (s Set) Maximum() lnwire.FeatureBit {
	switch s {
	case SetInvoice, SetInvoiceAmp:
		return lnwire.MaxBolt11Feature

	// The space available in other sets is > math.MaxUint16, so we just
	// return the maximum value our expression of a feature bit allows so
	// that any value will pass.
	default:
		return math.MaxUint16
	}
}