// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0

package auditv1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	v11 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
	protowire "google.golang.org/protobuf/encoding/protowire"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	hash "hash"
	math "math"
	sort "sort"
)

func cerbos_audit_v1_AccessLogEntry_hashpb_sum(m *AccessLogEntry, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.audit.v1.AccessLogEntry.call_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.CallId))

	}
	if _, ok := ignore["cerbos.audit.v1.AccessLogEntry.timestamp"]; !ok {
		if m.Timestamp != nil {
			google_protobuf_Timestamp_hashpb_sum(m.Timestamp, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.audit.v1.AccessLogEntry.peer"]; !ok {
		if m.Peer != nil {
			cerbos_audit_v1_Peer_hashpb_sum(m.Peer, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.audit.v1.AccessLogEntry.metadata"]; !ok {
		if len(m.Metadata) > 0 {
			keys := make([]string, len(m.Metadata))
			i := 0
			for k := range m.Metadata {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Metadata[k] != nil {
					cerbos_audit_v1_MetaValues_hashpb_sum(m.Metadata[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.audit.v1.AccessLogEntry.method"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Method))

	}
	if _, ok := ignore["cerbos.audit.v1.AccessLogEntry.status_code"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.StatusCode)))

	}
}

func cerbos_audit_v1_DecisionLogEntry_CheckResources_hashpb_sum(m *DecisionLogEntry_CheckResources, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.CheckResources.inputs"]; !ok {
		if len(m.Inputs) > 0 {
			for _, v := range m.Inputs {
				if v != nil {
					cerbos_engine_v1_CheckInput_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.CheckResources.outputs"]; !ok {
		if len(m.Outputs) > 0 {
			for _, v := range m.Outputs {
				if v != nil {
					cerbos_engine_v1_CheckOutput_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.CheckResources.error"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Error))

	}
}

func cerbos_audit_v1_DecisionLogEntry_PlanResources_hashpb_sum(m *DecisionLogEntry_PlanResources, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.PlanResources.input"]; !ok {
		if m.Input != nil {
			cerbos_engine_v1_PlanResourcesInput_hashpb_sum(m.Input, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.PlanResources.output"]; !ok {
		if m.Output != nil {
			cerbos_engine_v1_PlanResourcesOutput_hashpb_sum(m.Output, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.PlanResources.error"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Error))

	}
}

func cerbos_audit_v1_DecisionLogEntry_hashpb_sum(m *DecisionLogEntry, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.call_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.CallId))

	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.timestamp"]; !ok {
		if m.Timestamp != nil {
			google_protobuf_Timestamp_hashpb_sum(m.Timestamp, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.peer"]; !ok {
		if m.Peer != nil {
			cerbos_audit_v1_Peer_hashpb_sum(m.Peer, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.inputs"]; !ok {
		if len(m.Inputs) > 0 {
			for _, v := range m.Inputs {
				if v != nil {
					cerbos_engine_v1_CheckInput_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.outputs"]; !ok {
		if len(m.Outputs) > 0 {
			for _, v := range m.Outputs {
				if v != nil {
					cerbos_engine_v1_CheckOutput_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.error"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Error))

	}
	if m.Method != nil {
		if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.method"]; !ok {
			switch t := m.Method.(type) {
			case *DecisionLogEntry_CheckResources_:
				if t.CheckResources != nil {
					cerbos_audit_v1_DecisionLogEntry_CheckResources_hashpb_sum(t.CheckResources, hasher, ignore)
				}

			case *DecisionLogEntry_PlanResources_:
				if t.PlanResources != nil {
					cerbos_audit_v1_DecisionLogEntry_PlanResources_hashpb_sum(t.PlanResources, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.audit.v1.DecisionLogEntry.metadata"]; !ok {
		if len(m.Metadata) > 0 {
			keys := make([]string, len(m.Metadata))
			i := 0
			for k := range m.Metadata {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Metadata[k] != nil {
					cerbos_audit_v1_MetaValues_hashpb_sum(m.Metadata[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_audit_v1_MetaValues_hashpb_sum(m *MetaValues, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.audit.v1.MetaValues.values"]; !ok {
		if len(m.Values) > 0 {
			for _, v := range m.Values {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_audit_v1_Peer_hashpb_sum(m *Peer, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.audit.v1.Peer.address"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Address))

	}
	if _, ok := ignore["cerbos.audit.v1.Peer.auth_info"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.AuthInfo))

	}
	if _, ok := ignore["cerbos.audit.v1.Peer.user_agent"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.UserAgent))

	}
	if _, ok := ignore["cerbos.audit.v1.Peer.forwarded_for"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.ForwardedFor))

	}
}

func cerbos_engine_v1_AuxData_hashpb_sum(m *v1.AuxData, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.AuxData.jwt"]; !ok {
		if len(m.Jwt) > 0 {
			keys := make([]string, len(m.Jwt))
			i := 0
			for k := range m.Jwt {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Jwt[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Jwt[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_CheckInput_hashpb_sum(m *v1.CheckInput, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.CheckInput.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.resource"]; !ok {
		if m.Resource != nil {
			cerbos_engine_v1_Resource_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.CheckInput.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_engine_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
}

func cerbos_engine_v1_CheckOutput_ActionEffect_hashpb_sum(m *v1.CheckOutput_ActionEffect, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.ActionEffect.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Effect)))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.ActionEffect.policy"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Policy))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.ActionEffect.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_engine_v1_CheckOutput_hashpb_sum(m *v1.CheckOutput, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.resource_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.ResourceId))

	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.actions"]; !ok {
		if len(m.Actions) > 0 {
			keys := make([]string, len(m.Actions))
			i := 0
			for k := range m.Actions {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Actions[k] != nil {
					cerbos_engine_v1_CheckOutput_ActionEffect_hashpb_sum(m.Actions[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.effective_derived_roles"]; !ok {
		if len(m.EffectiveDerivedRoles) > 0 {
			for _, v := range m.EffectiveDerivedRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.CheckOutput.validation_errors"]; !ok {
		if len(m.ValidationErrors) > 0 {
			for _, v := range m.ValidationErrors {
				if v != nil {
					cerbos_schema_v1_ValidationError_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_PlanResourcesFilter_Expression_Operand_hashpb_sum(m *v1.PlanResourcesFilter_Expression_Operand, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Node != nil {
		if _, ok := ignore["cerbos.engine.v1.PlanResourcesFilter.Expression.Operand.node"]; !ok {
			switch t := m.Node.(type) {
			case *v1.PlanResourcesFilter_Expression_Operand_Value:
				if t.Value != nil {
					google_protobuf_Value_hashpb_sum(t.Value, hasher, ignore)
				}

			case *v1.PlanResourcesFilter_Expression_Operand_Expression:
				if t.Expression != nil {
					cerbos_engine_v1_PlanResourcesFilter_Expression_hashpb_sum(t.Expression, hasher, ignore)
				}

			case *v1.PlanResourcesFilter_Expression_Operand_Variable:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Variable))

			}
		}
	}
}

func cerbos_engine_v1_PlanResourcesFilter_Expression_hashpb_sum(m *v1.PlanResourcesFilter_Expression, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesFilter.Expression.operator"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Operator))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesFilter.Expression.operands"]; !ok {
		if len(m.Operands) > 0 {
			for _, v := range m.Operands {
				if v != nil {
					cerbos_engine_v1_PlanResourcesFilter_Expression_Operand_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_PlanResourcesFilter_hashpb_sum(m *v1.PlanResourcesFilter, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesFilter.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Kind)))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesFilter.condition"]; !ok {
		if m.Condition != nil {
			cerbos_engine_v1_PlanResourcesFilter_Expression_Operand_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
}

func cerbos_engine_v1_PlanResourcesInput_Resource_hashpb_sum(m *v1.PlanResourcesInput_Resource, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.Resource.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.Resource.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.Resource.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.Resource.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_engine_v1_PlanResourcesInput_hashpb_sum(m *v1.PlanResourcesInput, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.action"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Action))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.resource"]; !ok {
		if m.Resource != nil {
			cerbos_engine_v1_PlanResourcesInput_Resource_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_engine_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesInput.include_meta"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.IncludeMeta)))

	}
}

func cerbos_engine_v1_PlanResourcesOutput_hashpb_sum(m *v1.PlanResourcesOutput, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesOutput.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesOutput.action"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Action))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesOutput.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesOutput.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesOutput.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesOutput.filter"]; !ok {
		if m.Filter != nil {
			cerbos_engine_v1_PlanResourcesFilter_hashpb_sum(m.Filter, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesOutput.filter_debug"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.FilterDebug))

	}
	if _, ok := ignore["cerbos.engine.v1.PlanResourcesOutput.validation_errors"]; !ok {
		if len(m.ValidationErrors) > 0 {
			for _, v := range m.ValidationErrors {
				if v != nil {
					cerbos_schema_v1_ValidationError_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_engine_v1_Principal_hashpb_sum(m *v1.Principal, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Principal.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.engine.v1.Principal.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.Principal.roles"]; !ok {
		if len(m.Roles) > 0 {
			for _, v := range m.Roles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Principal.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Principal.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_engine_v1_Resource_hashpb_sum(m *v1.Resource, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Resource.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Resource.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_schema_v1_ValidationError_hashpb_sum(m *v11.ValidationError, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.schema.v1.ValidationError.path"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Path))

	}
	if _, ok := ignore["cerbos.schema.v1.ValidationError.message"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Message))

	}
	if _, ok := ignore["cerbos.schema.v1.ValidationError.source"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Source)))

	}
}

func google_protobuf_ListValue_hashpb_sum(m *structpb.ListValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.ListValue.values"]; !ok {
		if len(m.Values) > 0 {
			for _, v := range m.Values {
				if v != nil {
					google_protobuf_Value_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Struct_hashpb_sum(m *structpb.Struct, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Struct.fields"]; !ok {
		if len(m.Fields) > 0 {
			keys := make([]string, len(m.Fields))
			i := 0
			for k := range m.Fields {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Fields[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Fields[k], hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Timestamp_hashpb_sum(m *timestamppb.Timestamp, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Timestamp.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Timestamp.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}

func google_protobuf_Value_hashpb_sum(m *structpb.Value, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Kind != nil {
		if _, ok := ignore["google.protobuf.Value.kind"]; !ok {
			switch t := m.Kind.(type) {
			case *structpb.Value_NullValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.NullValue)))

			case *structpb.Value_NumberValue:
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(t.NumberValue)))

			case *structpb.Value_StringValue:
				_, _ = hasher.Write(protowire.AppendString(nil, t.StringValue))

			case *structpb.Value_BoolValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(t.BoolValue)))

			case *structpb.Value_StructValue:
				if t.StructValue != nil {
					google_protobuf_Struct_hashpb_sum(t.StructValue, hasher, ignore)
				}

			case *structpb.Value_ListValue:
				if t.ListValue != nil {
					google_protobuf_ListValue_hashpb_sum(t.ListValue, hasher, ignore)
				}

			}
		}
	}
}
