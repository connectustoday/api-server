import * as oauth2orize from 'oauth2orize';

let server;

function init(): void {
    server = oauth2orize.createServer();
    server.grant(oauth2orize.grant.code(function(client, requiredURI, user, ares, done) {
        let code = utils.uid(16);
        var ac = new AuthorizationCode(code, client.id, redirectURI, user.id, ares.scope);
        ac.save(function(err) {
            if (err) {
                return done(err);
            }
            return done(null, code);
        });
    }));
}
