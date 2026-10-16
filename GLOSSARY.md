# morn-go

A MongoDB library that maps one Go struct onto one collection and runs queries against it.

## Language

**Dao**:
A handle to one collection and the document template for that collection.
_Avoid_: repository, model, table

**Document**:
The struct named by a Dao's template.
_Avoid_: entity, record, model

**Clause**:
One query against a Dao's collection.
_Avoid_: builder, filter, statement

**Session**:
A transaction invoked through a callback. Work joins the transaction through the callback's context.
_Avoid_: transaction handle, unit of work

**GenID**:
The counter value handed out for a collection. The stored counter is the next value after the one just handed out.
_Avoid_: sequence, autoincrement, object id

**Query option**:
Settings for a query other than its filter, such as a hint or a collation. A limit, skip, or sort set on the Clause replaces the same setting on the query option.
_Avoid_: driver option, find option
