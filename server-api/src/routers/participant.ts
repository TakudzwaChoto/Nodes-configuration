import express, { Request, Response } from 'express';
import log4js from 'log4js';
import config from '../../config/constants.json';
import * as ParticipantService from '../fabricGateway/participant';

const logger = log4js.getLogger('user');
logger.level = config.logLevel;

const router = express.Router();

router.post('/add', async (req: Request, res: Response) => {
    try {
        const response = await ParticipantService.add(req);
        if (response.status === 'SUCCESS') {
            return res.status(200).send(response);
        } else if (response.status === 'DUPLICATE') {
            return res.status(409).send(response);
        } else if (response.status === 'ERROR') {
            return res.status(422).send(response);
        } else {
            return res.status(400).send(response);
        }
    } catch (error: any) {
        return res.status(401).send(error.toString());
    }
});

router.post('/search', async (req: Request, res: Response) => {
    try {
        const response = await ParticipantService.search(req);
        if (response.status === 'SUCCESS') {
            return res.status(200).send(response);
        } else if (response.status === 'ERROR') {
            return res.status(502).send(response);
        } else {
            return res.status(400).send(response);
        }
    } catch (error: any) {
        return res.status(401).send(error.toString());
    }
});

router.post('/update', async (req: Request, res: Response) => {
    try {
        const response = await ParticipantService.update(req);
        return res.status(200).send(response);
    } catch (error: any) {
        return res.status(401).send(error.toString());
    }
});

router.post('/getone', async (req: Request, res: Response) => {
    try {
        const response = await ParticipantService.getOne(req);
        return res.status(200).send(response);
    } catch (error: any) {
        return res.status(401).send(error.toString());
    }
});

router.post('/deleterecord', async (req: Request, res: Response) => {
    try {
        const response = await ParticipantService.deleteRecord(req);
        return res.status(200).send(response);
    } catch (error: any) {
        return res.status(401).send(error.toString());
    }
});

router.get('/getall', async (req: Request, res: Response) => {
    try {
        const response = await ParticipantService.getAll(req);
        return res.status(200).send(response);
    } catch (error: any) {
        return res.status(401).send(error.toString());
    }
});

router.post('/addprivate', async (req: Request, res: Response) => {
    try {
        const response = await ParticipantService.addPrivate(req);
        return res.status(200).send(response);
    } catch (error: any) {
        return res.status(401).send(error.toString());
    }
});

router.post('/getpvtdata', async (req: Request, res: Response) => {
    try {
        const response = await ParticipantService.getPvtData(req);
        return res.status(200).send(response);
    } catch (error: any) {
        return res.status(401).send(error.toString());
    }
});

router.post('/getipfsfile', async (req: Request, res: Response) => {
    try {
        const response = await ParticipantService.readIpfsFile(req);
        return res.status(200).send(response);
    } catch (error: any) {
        return res.status(401).send(error.toString());
    }
});

export default router;
