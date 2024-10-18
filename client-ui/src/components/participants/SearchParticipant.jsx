import React, { useEffect, useState } from 'react'
import ParticipantService from '../../services/participantApi'
import { ToastContainer, toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import ParticipantDetails from './ParticipantDetails';

const SearchParticipant = () => {
    const [searchValue, setSearchValue] = useState('');
    const [patients, setParticipant] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [showDetail, setShowDetail] = useState(false)
    const [participantDetails, setParticipantDetails] = useState()
    const navigate = useNavigate();
    const [searchMsg, setSearchMsg] = useState('')

    const handleSearch = async () => {
        setLoading(true);

        try {
            let resp = await ParticipantService.SearchParticipant(searchValue);

            if (resp.status === 200) {
                const ret_obj = JSON.parse(await resp.text())
                if (ret_obj?.objectBytes) {
                    setParticipant(JSON.parse(ret_obj.objectBytes))
                } else {
                    setParticipant([])
                    setSearchMsg('No participant record found.')
                }
                setParticipant(JSON.parse(ret_obj?.objectBytes) ? JSON.parse(ret_obj.objectBytes) : [])
            } else if (resp.status === 401) {
                setSearchMsg('There was some server error, please try again.')
            }
        } catch (err) {
            setSearchMsg(err.message);
        } finally {
            setLoading(false);
        }

    };

    const pullData = (mode) => {
        if (mode === 'search') {
            setShowDetail(false)
        }
    }

    const showDetails = (index) => {
        setShowDetail(true)
        setParticipantDetails(<ParticipantDetails participants={participants} index={index} fun={pullData} />)
    }

    return (
        <div className="flex flex-col min-h-[calc(100vh_-_120px)] pt-24 sm:pt-20 items-center">
            {!showDetail ?
                <>
                    <h1 className="text-3xl font-medium text-zinc-700 my-4">Search participant</h1>
                    <form onSubmit={setSearchValue}>
                        <div className="flex mb-4 px-4">
                            <input
                                type="text"
                                defaultValue={searchValue}
                                className="w-full text-slate-700 outline-none focus:bg-slate-50 border-gray-400 border-2 px-4 py-2 rounded-l-full"
                                placeholder="Enter participant name"
                                onChange={(e) => setSearchValue(e.target.value)}
                            />
                            <button
                                className={`bg-blue-500 text-white font-bold py-2 px-4 rounded-r-full transition-all duration-200 ${loading ? 'opacity-50 cursor-wait' : 'hover:bg-blue-600'
                                    }`}
                                onClick={handleSearch}
                                disabled={loading} >
                                {loading ? 'Loading...' : 'Search'}
                            </button>
                        </div>
                    </form>
                    <div className='flex max-w-[1240px] justify-center pt-2 w-full px-4'>
                        {error && <p className="text-red-500 mb-4">{error}</p>}
                        {patients?.length === 0 && !loading && <p className="text-red-500">{searchMsg}</p>}
                        {patients?.length > 0 && (
                            <table className="table-auto w-full">
                                <thead className='text-left'>
                                    <tr>
                                        <th className="px-4 py-2">First Name</th>
                                        <th className="px-4 py-2">Last Name</th>
                                        <th className="px-4 py-2">DOB</th>
                                        <th className="px-4 py-2 hidden sm:block">Phone</th>
                                    </tr>
                                </thead>
                                <tbody className='bg-teal-50'>
                                    {patients.map((participant, indx) => (
                                        (participant.participantID) !== '' ?
                                            <tr className='hover:cursor-pointer' key={indx} onClick={() => showDetails(indx)}>
                                                <td className="border px-4 py-1">{participant.fName}</td>
                                                <td className="border px-4 py-1">{participant.lName}</td>
                                                <td className="border px-4 py-1">{participant.dob}</td>
                                                <td className="border px-4 py-1 hidden sm:block">{participant.mobile}</td>
                                            </tr>
                                            : ''
                                    ))}
                                </tbody>
                            </table>
                        )}
                    </div>
                </>
                :
                patientDetails}

        </div>
    )
}

export default SearchPatient
