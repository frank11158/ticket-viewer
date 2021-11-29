import { Link } from 'react-router-dom'

const Ticket = ({ ticket }) => {

    return (
        <Link to={`ticket/${ticket.id}`}>
            <div className='ticket'>
                <div className='block status'>
                    <h4>{ticket.status.substring(0,1)}</h4>
                </div>
                <div className='block'>
                            <h4>{ticket.subject}</h4>
                </div>
            </div>
        </Link>
    )
}

export default Ticket
