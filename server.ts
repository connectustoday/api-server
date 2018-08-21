// ConnectUS Server
import app from "./app";

console.log("Starting ConnectUS API Server...");

const PORT = 3000;

app.listen(PORT, () => {
    console.log('Express server listening on port ' + PORT);
})