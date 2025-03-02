// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/pkg/flow/ent/inode"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirror"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/google/uuid"
)

// Mirror is the model entity for the Mirror schema.
type Mirror struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref string `json:"ref,omitempty"`
	// Cron holds the value of the "cron" field.
	Cron string `json:"cron,omitempty"`
	// PublicKey holds the value of the "public_key" field.
	PublicKey string `json:"publicKey,omitempty"`
	// PrivateKey holds the value of the "private_key" field.
	PrivateKey string `json:"private_key,omitempty"`
	// Passphrase holds the value of the "passphrase" field.
	Passphrase string `json:"passphrase,omitempty"`
	// Commit holds the value of the "commit" field.
	Commit string `json:"commit,omitempty"`
	// LastSync holds the value of the "last_sync" field.
	LastSync *time.Time `json:"last_sync,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MirrorQuery when eager-loading is set.
	Edges             MirrorEdges `json:"edges"`
	inode_mirror      *uuid.UUID
	namespace_mirrors *uuid.UUID
}

// MirrorEdges holds the relations/edges for other nodes in the graph.
type MirrorEdges struct {
	// Namespace holds the value of the namespace edge.
	Namespace *Namespace `json:"namespace,omitempty"`
	// Inode holds the value of the inode edge.
	Inode *Inode `json:"inode,omitempty"`
	// Activities holds the value of the activities edge.
	Activities []*MirrorActivity `json:"activities,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// NamespaceOrErr returns the Namespace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MirrorEdges) NamespaceOrErr() (*Namespace, error) {
	if e.loadedTypes[0] {
		if e.Namespace == nil {
			// The edge namespace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: namespace.Label}
		}
		return e.Namespace, nil
	}
	return nil, &NotLoadedError{edge: "namespace"}
}

// InodeOrErr returns the Inode value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MirrorEdges) InodeOrErr() (*Inode, error) {
	if e.loadedTypes[1] {
		if e.Inode == nil {
			// The edge inode was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: inode.Label}
		}
		return e.Inode, nil
	}
	return nil, &NotLoadedError{edge: "inode"}
}

// ActivitiesOrErr returns the Activities value or an error if the edge
// was not loaded in eager-loading.
func (e MirrorEdges) ActivitiesOrErr() ([]*MirrorActivity, error) {
	if e.loadedTypes[2] {
		return e.Activities, nil
	}
	return nil, &NotLoadedError{edge: "activities"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mirror) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mirror.FieldURL, mirror.FieldRef, mirror.FieldCron, mirror.FieldPublicKey, mirror.FieldPrivateKey, mirror.FieldPassphrase, mirror.FieldCommit:
			values[i] = new(sql.NullString)
		case mirror.FieldLastSync, mirror.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case mirror.FieldID:
			values[i] = new(uuid.UUID)
		case mirror.ForeignKeys[0]: // inode_mirror
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case mirror.ForeignKeys[1]: // namespace_mirrors
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Mirror", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mirror fields.
func (m *Mirror) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mirror.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case mirror.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				m.URL = value.String
			}
		case mirror.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				m.Ref = value.String
			}
		case mirror.FieldCron:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cron", values[i])
			} else if value.Valid {
				m.Cron = value.String
			}
		case mirror.FieldPublicKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_key", values[i])
			} else if value.Valid {
				m.PublicKey = value.String
			}
		case mirror.FieldPrivateKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_key", values[i])
			} else if value.Valid {
				m.PrivateKey = value.String
			}
		case mirror.FieldPassphrase:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field passphrase", values[i])
			} else if value.Valid {
				m.Passphrase = value.String
			}
		case mirror.FieldCommit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commit", values[i])
			} else if value.Valid {
				m.Commit = value.String
			}
		case mirror.FieldLastSync:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_sync", values[i])
			} else if value.Valid {
				m.LastSync = new(time.Time)
				*m.LastSync = value.Time
			}
		case mirror.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case mirror.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field inode_mirror", values[i])
			} else if value.Valid {
				m.inode_mirror = new(uuid.UUID)
				*m.inode_mirror = *value.S.(*uuid.UUID)
			}
		case mirror.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field namespace_mirrors", values[i])
			} else if value.Valid {
				m.namespace_mirrors = new(uuid.UUID)
				*m.namespace_mirrors = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryNamespace queries the "namespace" edge of the Mirror entity.
func (m *Mirror) QueryNamespace() *NamespaceQuery {
	return (&MirrorClient{config: m.config}).QueryNamespace(m)
}

// QueryInode queries the "inode" edge of the Mirror entity.
func (m *Mirror) QueryInode() *InodeQuery {
	return (&MirrorClient{config: m.config}).QueryInode(m)
}

// QueryActivities queries the "activities" edge of the Mirror entity.
func (m *Mirror) QueryActivities() *MirrorActivityQuery {
	return (&MirrorClient{config: m.config}).QueryActivities(m)
}

// Update returns a builder for updating this Mirror.
// Note that you need to call Mirror.Unwrap() before calling this method if this Mirror
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mirror) Update() *MirrorUpdateOne {
	return (&MirrorClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Mirror entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mirror) Unwrap() *Mirror {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mirror is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mirror) String() string {
	var builder strings.Builder
	builder.WriteString("Mirror(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", url=")
	builder.WriteString(m.URL)
	builder.WriteString(", ref=")
	builder.WriteString(m.Ref)
	builder.WriteString(", cron=")
	builder.WriteString(m.Cron)
	builder.WriteString(", public_key=")
	builder.WriteString(m.PublicKey)
	builder.WriteString(", private_key=")
	builder.WriteString(m.PrivateKey)
	builder.WriteString(", passphrase=")
	builder.WriteString(m.Passphrase)
	builder.WriteString(", commit=")
	builder.WriteString(m.Commit)
	if v := m.LastSync; v != nil {
		builder.WriteString(", last_sync=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Mirrors is a parsable slice of Mirror.
type Mirrors []*Mirror

func (m Mirrors) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
