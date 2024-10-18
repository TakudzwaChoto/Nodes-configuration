import React, { useEffect, useState } from 'react'
import { AiOutlineClose, AiOutlineMenu, AiOutlineLogin, AiOutlineLogout } from 'react-icons/ai'
import { BiSolidDownArrow } from 'react-icons/bi'
import { Link, useNavigate, useLocation } from 'react-router-dom'
import Login from './Login'
import { useSharedData } from '../context/StateContext'

const Navbar = () => {
    const { isLoggedIn, setIsLoggedIn } = useSharedData();
    const [nav, setNav] = useState(true)
    const [showModal, setShowModal] = useState(false);
    const [showParticipantMenu, setShowParticipantMenu] = useState(false);

    const navigate = useNavigate();
    const location = useLocation();

    const handleParticipantMenu = (item) => {
        switch (item) {
            case 'participant':
                setShowParticipantMenu(!showParticipantMenu)
                break;
            // Add other sub-menu items here..
            default:
                break;
        }
    }

    const handleNav = () => {
        setNav(!nav)
    }

    const handleLogInOut = () => {
        if (isLoggedIn) {
            localStorage.removeItem('user')
            setIsLoggedIn(false)
            navigate('/')
        } else {
            handleLoginClick()
        }
    }

    const handleLoginClick = () => {
        setShowModal(true);
    };

    const checkAuthentication = () => {
        if (!isLoggedIn) {
            handleLoginClick()
        }
    }

    return (
        <div className='fixed top-0 w-full content-center mx-auto h-20 md:h-14 shadow-md z-[99] text-gray-700 bg-[#fcfbfc] '>
            <div className='flex max-w-[1240px] justify-between mx-auto items-center w-full h-full pr-4 2xl:px-8'>
                <h1 className='w-full text-3xl uppercase drop-shadow-[0_2px_1px_rgba(0,0,0,0.3)] font-bold m-4 text-[#942e83]'>Water-qaulity-Management Network</h1>
                <div className='pr-6 hover:cursor-pointer' onClick={handleLogInOut}>
                    {isLoggedIn ?
                        <AiOutlineLogout size={22} color='red' markerWidth={20} strokeWidth={20} />
                        :
                        <AiOutlineLogin size={22} color='green' markerWidth={20} strokeWidth={20} />
                    }
                </div>
                <ul className='hidden md:flex whitespace-nowrap '>
                    <li className='p-4'><Link to='/'> Home</Link></li>
                    <li onMouseEnter={() => handleParticipantMenu('patient')} onMouseLeave={() => handleParticipantMenu('patient')} className="relative p-4 hover:cursor-pointer">
                        <div className='flex' onClick={checkAuthentication}>
                            <span>  Participants</span> <BiSolidDownArrow size={20} className='pl-2 mt-[2px]' />
                        </div>
                        {isLoggedIn && showParticipantMenu && (
                            <ul className="absolute w-[150px] top-11 left-0 bg-white py-2 rounded shadow-lg">
                                <li>
                                    <Link to="/addparticipant" className="block px-4 py-2 text-gray-900 hover:bg-gray-200">Add participant</Link>
                                </li>
                                <li>
                                    <Link to="/searchparticipant" className="block px-4 py-2 text-gray-900 hover:bg-gray-200">Search</Link>
                                </li>
                            </ul>
                        )}
                    </li>

                    <li className='p-4'><Link to="/services">Services</Link></li>
                    <li className='p-4'><Link to='/contactus'> Contact Us</Link></li>
                </ul>

                <div onClick={handleNav} className='block pr-4 md:hidden'>
                    {!nav ? <AiOutlineClose size={22} /> : <AiOutlineMenu size={22} />}
                </div>
                <div className='block md:hidden'>
                    <div className={!nav ? 'fixed left-0 top-0 h-full w-[65%] bg-[#fcfbfc] border-r border-gray-500 ease-in-out duration-500 ' :
                        'fixed left-[-110%] h-full top-0 ease-out duration-500'}>
                        <h1 className='w-full text-3xl font-bold m-4 uppercase text-[#942e83] drop-shadow-md '>Water Quality Management</h1>
                        <ul className='p-4'>
                            <li className='p-4 border-b border-gray-600' onClick={handleNav} ><Link to='/'> Home</Link></li>
                            <li onClick={() => handlePatientMenu('participant')} className="relative p-4 border-b border-gray-600">
                                <div className='flex'>
                                    <span> Participant</span><BiSolidDownArrow size={20} className='pl-2 mt-[2px]' />
                                </div>
                                {isLoggedIn && showParticipantMenu && (
                                    <ul className="absolute w-[150px] top-full left-0 bg-white py-2 rounded shadow-lg">
                                        <li onClick={handleNav}>
                                            <Link to="/addparticipant" className="block px-4 py-2 text-gray-900 hover:bg-gray-200">Add Participant</Link>
                                        </li>
                                        <li onClick={handleNav}>
                                            <Link to="/searchparticipant" className="block px-4 py-2 text-gray-900 hover:bg-gray-200">Search</Link>
                                        </li>

                                    </ul>
                                )}
                            </li>
                            <li onClick={handleNav} className='p-4 border-b border-gray-600'><Link to="/services">Services</Link></li>
                            <li onClick={handleNav} className='p-4 '><Link to='/contactus'> Contact Us</Link></li>
                        </ul>
                    </div>
                </div>

                {showModal && (
                    <div className="fixed z-10 inset-0 overflow-y-auto pr-4">
                        <div className="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                            <div className="fixed inset-0 transition-opacity" aria-hidden="true">
                                <div className="absolute inset-0 bg-gray-500 opacity-75"></div>
                            </div>

                            <div className="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl 
                            transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
                                <Login closeModal={() => setShowModal(false)} />
                            </div>
                        </div>
                    </div>
                )}
            </div>
        </div>
    )
}

export default Navbar