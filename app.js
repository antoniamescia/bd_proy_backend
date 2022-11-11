import express from 'express';
import helmet from 'helmet';
import cors from 'cors';
import bodyParser from 'body-parser';
import morgan from 'morgan';
import { getPeople } from './database.js';
const app = express();

app.use(helmet());
app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(morgan('dev'));

app.get('/people', async (req, res) => {
    const people = await getPeople();
    res.json(people);
});

app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).send('Client side error.');
});

app.listen(8080, () => {
    console.log('Server running on port 8080');
})
