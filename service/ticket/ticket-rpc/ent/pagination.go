// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attribute"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributegroup"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeoption"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/attributeset"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/entity"
	"github.com/anhao26/zero-cloud/service/ticket/ticket-rpc/ent/entityattribute"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type AttributePager struct {
	Order  attribute.OrderOption
	Filter func(*AttributeQuery) (*AttributeQuery, error)
}

// AttributePaginateOption enables pagination customization.
type AttributePaginateOption func(*AttributePager)

// DefaultAttributeOrder is the default ordering of Attribute.
var DefaultAttributeOrder = Desc(attribute.FieldID)

func newAttributePager(opts []AttributePaginateOption) (*AttributePager, error) {
	pager := &AttributePager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultAttributeOrder
	}
	return pager, nil
}

func (p *AttributePager) ApplyFilter(query *AttributeQuery) (*AttributeQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// AttributePageList is Attribute PageList result.
type AttributePageList struct {
	List        []*Attribute `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (a *AttributeQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...AttributePaginateOption,
) (*AttributePageList, error) {

	pager, err := newAttributePager(opts)
	if err != nil {
		return nil, err
	}

	if a, err = pager.ApplyFilter(a); err != nil {
		return nil, err
	}

	ret := &AttributePageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := a.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		a = a.Order(pager.Order)
	} else {
		a = a.Order(DefaultAttributeOrder)
	}

	a = a.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := a.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type AttributeGroupPager struct {
	Order  attributegroup.OrderOption
	Filter func(*AttributeGroupQuery) (*AttributeGroupQuery, error)
}

// AttributeGroupPaginateOption enables pagination customization.
type AttributeGroupPaginateOption func(*AttributeGroupPager)

// DefaultAttributeGroupOrder is the default ordering of AttributeGroup.
var DefaultAttributeGroupOrder = Desc(attributegroup.FieldID)

func newAttributeGroupPager(opts []AttributeGroupPaginateOption) (*AttributeGroupPager, error) {
	pager := &AttributeGroupPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultAttributeGroupOrder
	}
	return pager, nil
}

func (p *AttributeGroupPager) ApplyFilter(query *AttributeGroupQuery) (*AttributeGroupQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// AttributeGroupPageList is AttributeGroup PageList result.
type AttributeGroupPageList struct {
	List        []*AttributeGroup `json:"list"`
	PageDetails *PageDetails      `json:"pageDetails"`
}

func (ag *AttributeGroupQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...AttributeGroupPaginateOption,
) (*AttributeGroupPageList, error) {

	pager, err := newAttributeGroupPager(opts)
	if err != nil {
		return nil, err
	}

	if ag, err = pager.ApplyFilter(ag); err != nil {
		return nil, err
	}

	ret := &AttributeGroupPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := ag.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		ag = ag.Order(pager.Order)
	} else {
		ag = ag.Order(DefaultAttributeGroupOrder)
	}

	ag = ag.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := ag.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type AttributeOptionPager struct {
	Order  attributeoption.OrderOption
	Filter func(*AttributeOptionQuery) (*AttributeOptionQuery, error)
}

// AttributeOptionPaginateOption enables pagination customization.
type AttributeOptionPaginateOption func(*AttributeOptionPager)

// DefaultAttributeOptionOrder is the default ordering of AttributeOption.
var DefaultAttributeOptionOrder = Desc(attributeoption.FieldID)

func newAttributeOptionPager(opts []AttributeOptionPaginateOption) (*AttributeOptionPager, error) {
	pager := &AttributeOptionPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultAttributeOptionOrder
	}
	return pager, nil
}

func (p *AttributeOptionPager) ApplyFilter(query *AttributeOptionQuery) (*AttributeOptionQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// AttributeOptionPageList is AttributeOption PageList result.
type AttributeOptionPageList struct {
	List        []*AttributeOption `json:"list"`
	PageDetails *PageDetails       `json:"pageDetails"`
}

func (ao *AttributeOptionQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...AttributeOptionPaginateOption,
) (*AttributeOptionPageList, error) {

	pager, err := newAttributeOptionPager(opts)
	if err != nil {
		return nil, err
	}

	if ao, err = pager.ApplyFilter(ao); err != nil {
		return nil, err
	}

	ret := &AttributeOptionPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := ao.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		ao = ao.Order(pager.Order)
	} else {
		ao = ao.Order(DefaultAttributeOptionOrder)
	}

	ao = ao.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := ao.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type AttributeSetPager struct {
	Order  attributeset.OrderOption
	Filter func(*AttributeSetQuery) (*AttributeSetQuery, error)
}

// AttributeSetPaginateOption enables pagination customization.
type AttributeSetPaginateOption func(*AttributeSetPager)

// DefaultAttributeSetOrder is the default ordering of AttributeSet.
var DefaultAttributeSetOrder = Desc(attributeset.FieldID)

func newAttributeSetPager(opts []AttributeSetPaginateOption) (*AttributeSetPager, error) {
	pager := &AttributeSetPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultAttributeSetOrder
	}
	return pager, nil
}

func (p *AttributeSetPager) ApplyFilter(query *AttributeSetQuery) (*AttributeSetQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// AttributeSetPageList is AttributeSet PageList result.
type AttributeSetPageList struct {
	List        []*AttributeSet `json:"list"`
	PageDetails *PageDetails    `json:"pageDetails"`
}

func (as *AttributeSetQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...AttributeSetPaginateOption,
) (*AttributeSetPageList, error) {

	pager, err := newAttributeSetPager(opts)
	if err != nil {
		return nil, err
	}

	if as, err = pager.ApplyFilter(as); err != nil {
		return nil, err
	}

	ret := &AttributeSetPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := as.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		as = as.Order(pager.Order)
	} else {
		as = as.Order(DefaultAttributeSetOrder)
	}

	as = as.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := as.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type EntityPager struct {
	Order  entity.OrderOption
	Filter func(*EntityQuery) (*EntityQuery, error)
}

// EntityPaginateOption enables pagination customization.
type EntityPaginateOption func(*EntityPager)

// DefaultEntityOrder is the default ordering of Entity.
var DefaultEntityOrder = Desc(entity.FieldID)

func newEntityPager(opts []EntityPaginateOption) (*EntityPager, error) {
	pager := &EntityPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultEntityOrder
	}
	return pager, nil
}

func (p *EntityPager) ApplyFilter(query *EntityQuery) (*EntityQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// EntityPageList is Entity PageList result.
type EntityPageList struct {
	List        []*Entity    `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (e *EntityQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...EntityPaginateOption,
) (*EntityPageList, error) {

	pager, err := newEntityPager(opts)
	if err != nil {
		return nil, err
	}

	if e, err = pager.ApplyFilter(e); err != nil {
		return nil, err
	}

	ret := &EntityPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := e.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		e = e.Order(pager.Order)
	} else {
		e = e.Order(DefaultEntityOrder)
	}

	e = e.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := e.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type EntityAttributePager struct {
	Order  entityattribute.OrderOption
	Filter func(*EntityAttributeQuery) (*EntityAttributeQuery, error)
}

// EntityAttributePaginateOption enables pagination customization.
type EntityAttributePaginateOption func(*EntityAttributePager)

// DefaultEntityAttributeOrder is the default ordering of EntityAttribute.
var DefaultEntityAttributeOrder = Desc(entityattribute.FieldID)

func newEntityAttributePager(opts []EntityAttributePaginateOption) (*EntityAttributePager, error) {
	pager := &EntityAttributePager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultEntityAttributeOrder
	}
	return pager, nil
}

func (p *EntityAttributePager) ApplyFilter(query *EntityAttributeQuery) (*EntityAttributeQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// EntityAttributePageList is EntityAttribute PageList result.
type EntityAttributePageList struct {
	List        []*EntityAttribute `json:"list"`
	PageDetails *PageDetails       `json:"pageDetails"`
}

func (ea *EntityAttributeQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...EntityAttributePaginateOption,
) (*EntityAttributePageList, error) {

	pager, err := newEntityAttributePager(opts)
	if err != nil {
		return nil, err
	}

	if ea, err = pager.ApplyFilter(ea); err != nil {
		return nil, err
	}

	ret := &EntityAttributePageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := ea.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		ea = ea.Order(pager.Order)
	} else {
		ea = ea.Order(DefaultEntityAttributeOrder)
	}

	ea = ea.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := ea.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
