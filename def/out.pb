
�)
	def.protodefAgithub.com/envoyproxy/protoc-gen-validate/validate/validate.proto"�
RType2
from (	B���r2^[a-z][a-z0-9]{1,14}$Rfrom.
to (	B���r2^[a-z][a-z0-9]{1,14}$Rto2
verb (	B���r2^[a-z][a-z0-9]{1,14}$Rverb
multiple (Rmultiple,
	countType (2.def.CountTypeR	countType";
EType2
name (	B���r2^[a-z][a-z0-9]{1,14}$Rname"�
E2
Type (	B���r2^[a-z][a-z0-9]{1,14}$RType
ID (	RID
ID1 (	RID1
ID2 (	RID2
ID3 (	RID3
CTime (RCTime
UTime (RUTime 
State (2
.def.StateRState
Tags	 (	RTags$
Meta
 (2.def.E.MetaEntryRMeta-
Content (2.def.E.ContentEntryRContent
Score (RScore
Score1 (RScore1
	Resources (	R	Resources7
	MetaEntry
key (	Rkey
value (	Rvalue:8:
ContentEntry
key (	Rkey
value (	Rvalue:8"f
DefineETypeReq 
eType (2
.def.ETypeReType2
creationRTypes (2
.def.RTypeRcreationRTypes"J
CreateEWithRsReq
e (2.def.ERe 
related (2.def.ERrelated"u
RelationReq
from (2.def.ERfrom
to (2.def.ERto2
verb (	B���r2^[a-z][a-z0-9]{1,14}$Rverb"S
GetByIDsReq2
type (	B���r2^[a-z][a-z0-9]{1,14}$Rtype
ids (	Rids"#
EList
list (2.def.ERlist"
Empty"#
HasRelationResp
has (Rhas"�
Query
Field (	B	���rRField
Op (2.def.OpROp
Value (	RValue,
	ValueType (2.def.ValueTypeR	ValueType"1
Limit
From (	RFrom
Limit (RLimit"�
Update
Field (	B	���rRField)
Action (2.def.UpdateActionRAction
Value (	RValue,
	ValueType (2.def.ValueTypeR	ValueType"Y
Paged
List (2.def.ERList
HasMore (RHasMore
NextFrom (	RNextFrom"T
PagedIDs
List (	RList
HasMore (RHasMore
NextFrom (	RNextFrom"�
CountByState5
counts (2.def.CountByState.CountsEntryRcounts9
CountsEntry
key (	Rkey
value (Rvalue:8"�
Counts/
counts (2.def.Counts.CountsEntryRcountsL
CountsEntry
key (	Rkey'
value (2.def.CountByStateRvalue:8"c
HasRelationsReq
from (2.def.ERfrom
to (2.def.ERto
	relations (	R	relations"�
HasRelations>
	relations (2 .def.HasRelations.RelationsEntryR	relations<
RelationsEntry
key (	Rkey
value (Rvalue:8"j
GetRelationReq
from (2.def.ERfrom
relation (	Rrelation 
limit (2
.def.LimitRlimit"�
EX
entity (2.def.ERentity!
related (2.def.EXRrelated$
	resources (2.def.ER	resources#
counts (2.def.CountsRcounts5
hasRelations (2.def.HasRelationsRhasRelations1
children (2.def.EX.ChildrenEntryRchildrenI
ChildrenEntry
key (	Rkey"
value (2.def.EXPagedRvalue:8"\
EXPaged
List (2.def.EXRList
HasMore (RHasMore
NextFrom (	RNextFrom"%
EXList
List (2.def.EXRList"�
GetByQueryReq2
type (	B���r2^[a-z][a-z0-9]{1,14}$Rtype$
queries (2
.def.QueryRqueries3
sorts (2.def.GetByQueryReq.SortsEntryRsorts 
limit (2
.def.LimitRlimitF

SortsEntry
key (	Rkey"
value (2.def.SortDirRvalue:8"�
UpdateContentReq2
type (	B���r2^[a-z][a-z0-9]{1,14}$Rtype
id (	B	���rRid%
updates (2.def.UpdateRupdates"�
SetStateReq2
type (	B���r2^[a-z][a-z0-9]{1,14}$Rtype 
ids (	B���	�"rRids 
state (2
.def.StateRstate"a
	DeleteReq2
type (	B���r2^[a-z][a-z0-9]{1,14}$Rtype 
ids (	B���	�"rRids"]
GetCountsReq2
type (	B���r2^[a-z][a-z0-9]{1,14}$Rtype
id (	B	���rRid"�

SetMetaReq2
type (	B���r2^[a-z][a-z0-9]{1,14}$Rtype
id (	B	���rRid-
meta (2.def.SetMetaReq.MetaEntryRmeta7
	MetaEntry
key (	Rkey
value (	Rvalue:8"r
DeleteMetaReq2
type (	B���r2^[a-z][a-z0-9]{1,14}$Rtype
id (	B	���rRid
keys (	Rkeys"r
UpdateTagsReq2
type (	B���r2^[a-z][a-z0-9]{1,14}$Rtype
id (	B	���rRid
tags (	Rtags*K
	CountType
	CountBoth 
CountFromEnd

CountToEnd
	CountNone*q
State
StateDeleted 
StatePrivate

State1

State2

State3
StateFriend
StatePublic*<
Op
Eq 
Gt
Ge
Lt
Le
Ne
In*6
UpdateAction
Set 
Incr
Add

Remove*L
	ValueType

String 	
Int64
Int

Double
Bool	
Bytes*
SortDir
Asc 
Desc2�
Storage.
defineEType.def.DefineETypeReq
.def.Empty%
defineRType
.def.RType
.def.Empty'
create.def.CreateEWithRsReq.def.E.
createRelation.def.RelationReq
.def.Empty.
removeRelation.def.RelationReq
.def.Empty7
hasRelations.def.HasRelationsReq.def.HasRelations1
getRelation.def.GetRelationReq.def.PagedIDs(
getByIds.def.GetByIDsReq
.def.EList,

getByQuery.def.GetByQueryReq
.def.Paged2
updateContent.def.UpdateContentReq
.def.Empty(
setState.def.SetStateReq
.def.Empty$
delete.def.DeleteReq
.def.Empty&
setMeta.def.SetMetaReq
.def.Empty,

deleteMeta.def.DeleteMetaReq
.def.Empty)
addTags.def.UpdateTagsReq
.def.Empty,

deleteTags.def.UpdateTagsReq
.def.Empty+
	getCounts.def.GetCountsReq.def.Countsbproto3
�
	dsl.protodef	def.proto"�
PopulationOption<
related (2".def.PopulationOption.RelatedEntryRrelated

withCounts (R
withCounts$
withResources (RwithResources2
hasRelationsWith (2.def.ERhasRelationsWith&
hasRelationsOf (	RhasRelationsOf:
RelatedEntry
key (Rkey
value (	Rvalue:8"�
EntityQuery
qt (2.def.QueryTypeRqt
type (	Rtype
id (	Rid
ids (	Rids+
popOp (2.def.PopulationOptionRpopOp
user (2.def.EXRuser
action (	Raction$
queries (2
.def.QueryRqueries1
sorts	 (2.def.EntityQuery.SortsEntryRsorts
fromID
 (	RfromID
	withLimit (R	withLimitF

SortsEntry
key (	Rkey"
value (2.def.SortDirRvalue:8"�
EntityCreate
type (	Rtype
e_ (2.def.ERe"
related_ (2.def.EXRrelated%

resources_ (2.def.ER	resources
user (2.def.EXRuser
action (	Raction"�
EntityUpdate
type (	Rtype
ofID (	RofID$
queries (2
.def.QueryRqueries%
updates (2.def.UpdateRupdates
user (2.def.EXRuser
action (	Raction"�
RelationQuery
type (	Rtype
ofID (2.def.EXRofID
relation (	Rrelation
user (2.def.EXRuser+
popOp (2.def.PopulationOptionRpopOp
fromID (	RfromID
	withLimit (R	withLimit"�
RelationCreate
from_ (2.def.EXRfrom
to_ (2.def.EXRto
verb_ (	Rverb
user (2.def.EXRuser
isAdd_ (RisAdd*+
	QueryType
One 	
ByIds

Paging2�
DSL 
one.def.EntityQuery.def.EX$
ids.def.EntityQuery.def.EXList'
paged.def.EntityQuery.def.EXPaged*
createEntity.def.EntityCreate.def.EX*
updateEntity.def.EntityUpdate.def.EX1
queryRelation.def.RelationQuery.def.EXPaged1
createRelation.def.RelationCreate
.def.Emptybproto3