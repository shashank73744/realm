// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"

	"github.com/kcarretto/realm/ent/credential"
	"github.com/kcarretto/realm/ent/predicate"
	"github.com/kcarretto/realm/ent/target"
)

// CredentialWhereInput represents a where input for filtering Credential queries.
type CredentialWhereInput struct {
	Not *CredentialWhereInput   `json:"not,omitempty"`
	Or  []*CredentialWhereInput `json:"or,omitempty"`
	And []*CredentialWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "principal" field predicates.
	Principal             *string  `json:"principal,omitempty"`
	PrincipalNEQ          *string  `json:"principalNEQ,omitempty"`
	PrincipalIn           []string `json:"principalIn,omitempty"`
	PrincipalNotIn        []string `json:"principalNotIn,omitempty"`
	PrincipalGT           *string  `json:"principalGT,omitempty"`
	PrincipalGTE          *string  `json:"principalGTE,omitempty"`
	PrincipalLT           *string  `json:"principalLT,omitempty"`
	PrincipalLTE          *string  `json:"principalLTE,omitempty"`
	PrincipalContains     *string  `json:"principalContains,omitempty"`
	PrincipalHasPrefix    *string  `json:"principalHasPrefix,omitempty"`
	PrincipalHasSuffix    *string  `json:"principalHasSuffix,omitempty"`
	PrincipalEqualFold    *string  `json:"principalEqualFold,omitempty"`
	PrincipalContainsFold *string  `json:"principalContainsFold,omitempty"`

	// "secret" field predicates.
	Secret             *string  `json:"secret,omitempty"`
	SecretNEQ          *string  `json:"secretNEQ,omitempty"`
	SecretIn           []string `json:"secretIn,omitempty"`
	SecretNotIn        []string `json:"secretNotIn,omitempty"`
	SecretGT           *string  `json:"secretGT,omitempty"`
	SecretGTE          *string  `json:"secretGTE,omitempty"`
	SecretLT           *string  `json:"secretLT,omitempty"`
	SecretLTE          *string  `json:"secretLTE,omitempty"`
	SecretContains     *string  `json:"secretContains,omitempty"`
	SecretHasPrefix    *string  `json:"secretHasPrefix,omitempty"`
	SecretHasSuffix    *string  `json:"secretHasSuffix,omitempty"`
	SecretEqualFold    *string  `json:"secretEqualFold,omitempty"`
	SecretContainsFold *string  `json:"secretContainsFold,omitempty"`

	// "kind" field predicates.
	Kind      *credential.Kind  `json:"kind,omitempty"`
	KindNEQ   *credential.Kind  `json:"kindNEQ,omitempty"`
	KindIn    []credential.Kind `json:"kindIn,omitempty"`
	KindNotIn []credential.Kind `json:"kindNotIn,omitempty"`
}

// Filter applies the CredentialWhereInput filter on the CredentialQuery builder.
func (i *CredentialWhereInput) Filter(q *CredentialQuery) (*CredentialQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering credentials.
// An error is returned if the input is empty or invalid.
func (i *CredentialWhereInput) P() (predicate.Credential, error) {
	var predicates []predicate.Credential
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, credential.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Credential, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, credential.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Credential, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, credential.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, credential.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, credential.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, credential.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, credential.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, credential.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, credential.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, credential.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, credential.IDLTE(*i.IDLTE))
	}
	if i.Principal != nil {
		predicates = append(predicates, credential.PrincipalEQ(*i.Principal))
	}
	if i.PrincipalNEQ != nil {
		predicates = append(predicates, credential.PrincipalNEQ(*i.PrincipalNEQ))
	}
	if len(i.PrincipalIn) > 0 {
		predicates = append(predicates, credential.PrincipalIn(i.PrincipalIn...))
	}
	if len(i.PrincipalNotIn) > 0 {
		predicates = append(predicates, credential.PrincipalNotIn(i.PrincipalNotIn...))
	}
	if i.PrincipalGT != nil {
		predicates = append(predicates, credential.PrincipalGT(*i.PrincipalGT))
	}
	if i.PrincipalGTE != nil {
		predicates = append(predicates, credential.PrincipalGTE(*i.PrincipalGTE))
	}
	if i.PrincipalLT != nil {
		predicates = append(predicates, credential.PrincipalLT(*i.PrincipalLT))
	}
	if i.PrincipalLTE != nil {
		predicates = append(predicates, credential.PrincipalLTE(*i.PrincipalLTE))
	}
	if i.PrincipalContains != nil {
		predicates = append(predicates, credential.PrincipalContains(*i.PrincipalContains))
	}
	if i.PrincipalHasPrefix != nil {
		predicates = append(predicates, credential.PrincipalHasPrefix(*i.PrincipalHasPrefix))
	}
	if i.PrincipalHasSuffix != nil {
		predicates = append(predicates, credential.PrincipalHasSuffix(*i.PrincipalHasSuffix))
	}
	if i.PrincipalEqualFold != nil {
		predicates = append(predicates, credential.PrincipalEqualFold(*i.PrincipalEqualFold))
	}
	if i.PrincipalContainsFold != nil {
		predicates = append(predicates, credential.PrincipalContainsFold(*i.PrincipalContainsFold))
	}
	if i.Secret != nil {
		predicates = append(predicates, credential.SecretEQ(*i.Secret))
	}
	if i.SecretNEQ != nil {
		predicates = append(predicates, credential.SecretNEQ(*i.SecretNEQ))
	}
	if len(i.SecretIn) > 0 {
		predicates = append(predicates, credential.SecretIn(i.SecretIn...))
	}
	if len(i.SecretNotIn) > 0 {
		predicates = append(predicates, credential.SecretNotIn(i.SecretNotIn...))
	}
	if i.SecretGT != nil {
		predicates = append(predicates, credential.SecretGT(*i.SecretGT))
	}
	if i.SecretGTE != nil {
		predicates = append(predicates, credential.SecretGTE(*i.SecretGTE))
	}
	if i.SecretLT != nil {
		predicates = append(predicates, credential.SecretLT(*i.SecretLT))
	}
	if i.SecretLTE != nil {
		predicates = append(predicates, credential.SecretLTE(*i.SecretLTE))
	}
	if i.SecretContains != nil {
		predicates = append(predicates, credential.SecretContains(*i.SecretContains))
	}
	if i.SecretHasPrefix != nil {
		predicates = append(predicates, credential.SecretHasPrefix(*i.SecretHasPrefix))
	}
	if i.SecretHasSuffix != nil {
		predicates = append(predicates, credential.SecretHasSuffix(*i.SecretHasSuffix))
	}
	if i.SecretEqualFold != nil {
		predicates = append(predicates, credential.SecretEqualFold(*i.SecretEqualFold))
	}
	if i.SecretContainsFold != nil {
		predicates = append(predicates, credential.SecretContainsFold(*i.SecretContainsFold))
	}
	if i.Kind != nil {
		predicates = append(predicates, credential.KindEQ(*i.Kind))
	}
	if i.KindNEQ != nil {
		predicates = append(predicates, credential.KindNEQ(*i.KindNEQ))
	}
	if len(i.KindIn) > 0 {
		predicates = append(predicates, credential.KindIn(i.KindIn...))
	}
	if len(i.KindNotIn) > 0 {
		predicates = append(predicates, credential.KindNotIn(i.KindNotIn...))
	}

	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/kcarretto/realm/ent: empty predicate CredentialWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return credential.And(predicates...), nil
	}
}

// TargetWhereInput represents a where input for filtering Target queries.
type TargetWhereInput struct {
	Not *TargetWhereInput   `json:"not,omitempty"`
	Or  []*TargetWhereInput `json:"or,omitempty"`
	And []*TargetWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "forwardConnectIP" field predicates.
	ForwardConnectIP             *string  `json:"forwardconnectip,omitempty"`
	ForwardConnectIPNEQ          *string  `json:"forwardconnectipNEQ,omitempty"`
	ForwardConnectIPIn           []string `json:"forwardconnectipIn,omitempty"`
	ForwardConnectIPNotIn        []string `json:"forwardconnectipNotIn,omitempty"`
	ForwardConnectIPGT           *string  `json:"forwardconnectipGT,omitempty"`
	ForwardConnectIPGTE          *string  `json:"forwardconnectipGTE,omitempty"`
	ForwardConnectIPLT           *string  `json:"forwardconnectipLT,omitempty"`
	ForwardConnectIPLTE          *string  `json:"forwardconnectipLTE,omitempty"`
	ForwardConnectIPContains     *string  `json:"forwardconnectipContains,omitempty"`
	ForwardConnectIPHasPrefix    *string  `json:"forwardconnectipHasPrefix,omitempty"`
	ForwardConnectIPHasSuffix    *string  `json:"forwardconnectipHasSuffix,omitempty"`
	ForwardConnectIPEqualFold    *string  `json:"forwardconnectipEqualFold,omitempty"`
	ForwardConnectIPContainsFold *string  `json:"forwardconnectipContainsFold,omitempty"`
}

// Filter applies the TargetWhereInput filter on the TargetQuery builder.
func (i *TargetWhereInput) Filter(q *TargetQuery) (*TargetQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering targets.
// An error is returned if the input is empty or invalid.
func (i *TargetWhereInput) P() (predicate.Target, error) {
	var predicates []predicate.Target
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, target.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Target, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, target.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Target, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, target.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, target.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, target.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, target.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, target.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, target.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, target.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, target.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, target.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, target.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, target.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, target.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, target.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, target.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, target.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, target.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, target.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, target.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, target.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, target.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, target.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, target.NameContainsFold(*i.NameContainsFold))
	}
	if i.ForwardConnectIP != nil {
		predicates = append(predicates, target.ForwardConnectIPEQ(*i.ForwardConnectIP))
	}
	if i.ForwardConnectIPNEQ != nil {
		predicates = append(predicates, target.ForwardConnectIPNEQ(*i.ForwardConnectIPNEQ))
	}
	if len(i.ForwardConnectIPIn) > 0 {
		predicates = append(predicates, target.ForwardConnectIPIn(i.ForwardConnectIPIn...))
	}
	if len(i.ForwardConnectIPNotIn) > 0 {
		predicates = append(predicates, target.ForwardConnectIPNotIn(i.ForwardConnectIPNotIn...))
	}
	if i.ForwardConnectIPGT != nil {
		predicates = append(predicates, target.ForwardConnectIPGT(*i.ForwardConnectIPGT))
	}
	if i.ForwardConnectIPGTE != nil {
		predicates = append(predicates, target.ForwardConnectIPGTE(*i.ForwardConnectIPGTE))
	}
	if i.ForwardConnectIPLT != nil {
		predicates = append(predicates, target.ForwardConnectIPLT(*i.ForwardConnectIPLT))
	}
	if i.ForwardConnectIPLTE != nil {
		predicates = append(predicates, target.ForwardConnectIPLTE(*i.ForwardConnectIPLTE))
	}
	if i.ForwardConnectIPContains != nil {
		predicates = append(predicates, target.ForwardConnectIPContains(*i.ForwardConnectIPContains))
	}
	if i.ForwardConnectIPHasPrefix != nil {
		predicates = append(predicates, target.ForwardConnectIPHasPrefix(*i.ForwardConnectIPHasPrefix))
	}
	if i.ForwardConnectIPHasSuffix != nil {
		predicates = append(predicates, target.ForwardConnectIPHasSuffix(*i.ForwardConnectIPHasSuffix))
	}
	if i.ForwardConnectIPEqualFold != nil {
		predicates = append(predicates, target.ForwardConnectIPEqualFold(*i.ForwardConnectIPEqualFold))
	}
	if i.ForwardConnectIPContainsFold != nil {
		predicates = append(predicates, target.ForwardConnectIPContainsFold(*i.ForwardConnectIPContainsFold))
	}

	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/kcarretto/realm/ent: empty predicate TargetWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return target.And(predicates...), nil
	}
}
