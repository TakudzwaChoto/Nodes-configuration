import {
    API_ADD_PARTICIPANT, API_UPDATE_PARTICIPANT, API_SEARCH_PARTICIPANT, HTTP_HEADER, API_GET_ALL_PARTICIPANTS,
    API_GET_PATI_DETAILS, API_ADD_PATI_PRIVATE_DATA, API_GET_PATI_PVT_DATA, HTTP_HEADER_FORMDATA,
    API_GET_PATI_IPFS_FILE, API_DELETE_PATI
} from './Constants.js';

const AddParticipant = async (participant) => {
    try {
        const resp = await fetch(API_ADD_PARTICIPANT, {
            headers: HTTP_HEADER(),
            method: 'POST',
            body: JSON.stringify(participant),
        })
        return (resp)
    } catch (error) {
        throw error;
    }
}

const UpdateParticipant = async (participant) => {
    try {
        const resp = await fetch(API_UPDATE_PARTICIPANT, {
            headers: HTTP_HEADER(),
            method: 'POST',
            body: JSON.stringify(participant),
        })
        return (resp)
    } catch (error) {
        throw error;
    }
}


const SearchParticipant = async (searchString) => {
    try {
        const resp = await fetch(API_SEARCH_PARTICIPANT, {
            headers: HTTP_HEADER(),
            method: 'POST',
            body: JSON.stringify({ searchString }),
        })

        return resp

    } catch (error) {
        throw error;
    }
}

const AddPvtWaterDetails = async (privateData) => {
    try {
        const resp = await fetch(API_ADD_PATI_PRIVATE_DATA, {
            headers: HTTP_HEADER_FORMDATA(),
            method: 'POST',
            body: privateData,
        })

        return resp

    } catch (error) {
        throw error;
    }
}

const GetPvtWaterRecords = async (participantID) => {
    console.log(participantID)
    try {
        const resp = await fetch(API_GET_PATI_PVT_DATA, {
            headers: HTTP_HEADER(),
            method: 'POST',
            body: JSON.stringify({ participantID }),
        })
        return resp
    } catch (error) {
        throw error;
    }
}

const GetParticipantDetails = async (participantID) => {
    try {
        const resp = await fetch(API_GET_PATI_DETAILS, {
            headers: HTTP_HEADER(),
            method: 'POST',
            body: JSON.stringify({ participantID }),
        })

        return resp
    } catch (error) {
        return error;
    }
}

const GetParticipantList = async () => {
    try {
        const resp = await fetch(API_GET_ALL_PARTICIPANT, {
            headers: HTTP_HEADER(),
            method: 'GET',
        })

        return resp
    } catch (error) {
        throw error;
    }
}

const DeleteParticipant = async (participantID) => {
    try {
        const resp = await fetch(API_DELETE_PATI, {
            headers: HTTP_HEADER(),
            method: 'POST',
            body: JSON.stringify({ participantID }),
        })
        const empstr = await resp.text()
        return empstr
    } catch (error) {
        return error;
    }
}

const GetIpfsFile = async (props) => {
    try {
        const resp = await fetch(API_GET_PATI_IPFS_FILE, {
            headers: HTTP_HEADER(),
            method: 'POST',
            body: JSON.stringify({ props }),

        })
        var buffer = await resp.arrayBuffer();
        return buffer;
    } catch (error) {
        return error;
    }
}


export default {
    AddParticipant,
    UpdateParticipant,
    SearchParticipant,
    GetParticipantList,
    GetParticipantDetails,
    DeleteParticipant,
    AddPvtWaterDetails,
    GetPvtWaterRecords,
    GetIpfsFile
}