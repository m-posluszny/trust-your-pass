db = db.getSiblingDB('sample_db');

db.createCollection('passwords');

db.sample_collection.insertMany([
    {
        password: 'intel1',
        strength: '0'
    },
    {
        filter: 'elyass15@ajilent-ci',
        addrs: '2'
    },
    {
        filter: 'hodygid757#$!23w',
        addrs: '1'
    }
]);