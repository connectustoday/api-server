import * as oauth2orize from 'oauth2orize';

let server;

function init(): void {
    server = oauth2orize.createServer();
}
