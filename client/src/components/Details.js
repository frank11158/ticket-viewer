import { useState, useEffect } from "react"
import { Link, useParams } from 'react-router-dom'

const Details = ({ tickets, setAPIErr }) => {
    const [ticket, setTicket] = useState()
    const [userInfo, setUserInfo] = useState('')
    const {id} = useParams()

    useEffect(() => {
        const getTicket = async () => {
            const ticket = tickets.find((ticket) => (
                ticket.id === parseInt({id}.id)
            ))
            setTicket(ticket)
        }
        const getUserInfo = async () => {
            if (ticket !== undefined) {
                try {
                    const response = await fetchUserInfo(ticket.requester_id)
                    if (response.msg !== "Ok") {
                        if (response.data !== null) {
                            setAPIErr(response.data.error)
                        } else {
                            setAPIErr("Something went wrong!")
                        }
                        throw new Error("Failed to fetch user info")
                    }
                    setUserInfo(response.data.user.name)
                } catch (e) {
                    console.log(e.error)
                }
            }
        }

        getTicket()
        getUserInfo()
    }, [id, ticket, tickets])

    // Fetch user info
    const fetchUserInfo = async (userID) => {
        const res = await fetch(`http://localhost:25976/api/v1/users/${userID}`)
        const data = await res.json()
        
        return data
    }

    return (
        <div>
            {
                ticket && 
                <>
                    <div className='subject'>
                        <h2>{ticket.subject}</h2>
                        <h5>Requested by {userInfo} on {new Date(ticket.updated_at).toLocaleDateString("en-US")}</h5>
                    </div>
                    <hr></hr>
                    <div className='description'>
                        <p>{ticket.description}</p>
                    </div>
                </>
            }
            <Link to='/'>Go Back</Link>
        </div>
    )
}

export default Details
