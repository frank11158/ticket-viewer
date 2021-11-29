import { useState, useEffect } from "react"
import Snackbar from "@material-ui/core/Snackbar"
import Alert from '@material-ui/core/Alert'

const APIErrorSnackBar = ({ open, message, setAPIErr }) => {
    const [op, setOpen] = useState(false)
    const [msg, setMsg] = useState("")

    useEffect(() => {
        if (open) {
            setOpen(open)
            setMsg(message)
        }
    }, [open, message])

    const handleClose = () => {
        setOpen(false)
        setMsg("")
        setAPIErr("")
    }

    return (
        <div>
            <Snackbar
                open={op}
                autoHideDuration={5000}
                message={msg}
            >
                <Alert onClose={handleClose} severity="error" sx={{ width: '100%' }}>
                    {msg}
                </Alert>
            </Snackbar>
        </div>
    )
}

export default APIErrorSnackBar
