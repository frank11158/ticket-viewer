import Ticket from "./Ticket"

const Tickets = ({ tickets  }) => {
    return (
        <>
            {tickets.map((ticket) => (
                <Ticket key={ticket.id} ticket={ticket} />
            ))}  
        </>
    )
}

export default Tickets
