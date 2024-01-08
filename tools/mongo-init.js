db = db.getSiblingDB('sample_db');

db.createCollection('passwords');

db.passwords.insertMany([
    {
        password: 'intel1',
        strength: 0,
        IsProcessing: true
    },
    {
        password: 'elyass15@ajilent-ci',
        strength: 2,
        IsProcessing: true
    },
    {
        password: 'hodygid757#$!23w',
        strength: 1,
        IsProcessing: true
    },
    {
        password: 'notEstimatedYet',
        strength: -1,
        IsProcessing: false
    }
]);