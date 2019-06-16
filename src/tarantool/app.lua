#!/usr/bin/env tarantool

-- Настроить базу данных
box.cfg {
    listen = 3301
}

console = require("console")
console.delimiter("!")

b = box.schema.space.create('blocks')!

box.schema.user.grant("guest","read,write,execute,create,drop","universe")!

b:format({
{name = "ID", type = "unsigned"},
{name = "DocumnetID", type = "unsigned"},
{name = "Data", type = "string"},
{name = "Line", type = "unsigned"},
{name = "Column", type = "unsigned"}
})!

b:create_index('ID', {
    type = "hash",
    parts = {'ID'}
})!

-- function test_data(data)
--     print('data: ', data)
--     return  data
-- end

-- msgpack = require('msgpack')