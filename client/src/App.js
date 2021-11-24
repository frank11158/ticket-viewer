import { useState, useEffect } from "react"
import Tickets from "./components/Tickets"

function App() {
  const [tickets, setTickets] = useState([])

  useEffect(() => {
    const getTickets = async () => {
      const ticketsFromServer = await fetchTickets()
      setTickets(ticketsFromServer)
    }

    getTickets()
  }, [])

  // Fetch Tickets
  const fetchTickets = async () => {
    const res = await fetch('http://localhost:5000/tickets')
    const data = await res.json()
    
    return data
  }

  return (
    <div className="container">
      <>
        {tickets.length > 0 ? <Tickets tickets={tickets} /> : 'No Ticket to Show'}
      </>
    </div>
  );
}

export default App;
