
const Ticket = ({ ticket }) => {
    return (
        <div>
            <h3>
                {ticket.text}
                {/* <FaTimes style={{color: 'red', cursor: 'pointer'}} onClick={() => onDelete(task.id)}/> */}
            </h3>
        </div>
    )
}

export default Ticket
