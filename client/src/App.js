import { useState, useEffect } from "react"
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Header from "./components/Header"
import Tickets from "./components/Tickets"
import Details from "./components/Details"
import Pagination from "./components/Pagination"
import APIErrorSnackBar from "./components/APIErrorSnackBar"

function App() {
  const [tickets, setTickets] = useState([])
  const [totalTickets, setTotalTickets] = useState(0)
  const [currentPage, setCurrentPage] = useState(1)
  const [ticketsPerPage] = useState(25)
  const [APIerr, setAPIErr] = useState("")

  useEffect(() => {
    const getTickets = async (currentPage, perPage) => {
      try {
        const response = await fetchTickets(currentPage, perPage)
        if (response.msg !== "Ok") {
          if (response.data !== null) {
            setAPIErr(response.data.error)
          } else {
            setAPIErr("Something went wrong!")
          }
          throw new Error("Failed to fetch tickets")
        }
        setTotalTickets(response.data.count)
        setTickets(response.data.tickets)
      } catch (e) {
        console.log(e.error)
      }
    }

    getTickets(currentPage, ticketsPerPage)
  }, [currentPage, ticketsPerPage])

  // Fetch Tickets
  const fetchTickets = async (currentPage, perPage) => {
    const res = await fetch(`http://localhost:25976/api/v1/tickets?page=${currentPage}&per_page=${perPage}`)
    const data = await res.json()
    
    return data
  }

  // Change page
  const paginate = pageNumber => setCurrentPage(pageNumber);


  return (
    <div className="container">
      <Router>
        <Header />
        <Routes>
          <Route path='/' exact element={
            <>
              {tickets && tickets.length > 0 ? <Tickets tickets={tickets} /> : 'Loading...'}
              <Pagination ticketsPerPage={ticketsPerPage} totalTickets={totalTickets} paginate={paginate}/>
            </>
          }/>
          <Route path='/ticket/:id' element={<Details tickets={tickets} setAPIErr={setAPIErr} />} />
        </Routes>
        <APIErrorSnackBar open={!!APIerr} message={APIerr} setAPIErr={setAPIErr}/>
      </Router>
    </div>
  )
}

export default App;