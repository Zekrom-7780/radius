import express from 'express';
import path from 'path';
import * as routes from "./routes";
import { SecretClient } from "@azure/keyvault-secrets";
import { ManagedIdentityCredential } from "@azure/identity";

export async function main(): Promise<void> {
  const app = express();
  const port = process.env.PORT || 3000;

  // Using the KV_URI env-var to get the KeyVault URI
  const kvURI = process.env.KV_URI || '';
  var connectionString:string = '';

  if (kvURI) {
    // Using the DBCONNECTION secret to get the secret name for the DB connection string
    const secretName = "DBCONNECTION";
    console.log("Secret name: ", secretName)

    // Access the key vault to fetch the DB connection string
    const credential = new ManagedIdentityCredential();
    const client = new SecretClient(kvURI, credential);
    connectionString = await (await client.getSecret(secretName)).value || '';
    console.log("Retrieved DB connection string from Key Vault")
  } else {
    connectionString = process.env.DBCONNECTION || ''
  }

  if (connectionString) {
    app.set("connectionString", connectionString);
  }

  app.use(express.json());

  app.set("views", path.join(__dirname, "views"));
  app.set("view engine", "ejs");

  app.use(express.static(path.join(__dirname, "www")));

  routes.register(app);

  function logError(err: any, req: any, res: any, next: any) {
    console.log(err)
    next()
  }
  app.use(logError)

  process.on('SIGINT', function() {
    console.log( "\nGracefully shutting down from SIGINT (Ctrl-C)" );
    process.exit(1);
  });

  app.listen(port, () =>
    console.log(`App listening on port ${port}!`),
  );
}

main()
