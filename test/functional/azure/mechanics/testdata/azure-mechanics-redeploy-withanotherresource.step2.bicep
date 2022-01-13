resource app 'radius.dev/Application@v1alpha3' = {
  name: 'azure-mechanics-redeploy-withanotherresource'

  resource a 'Container' = {
    name: 'a'
    properties: {
      container: {
        image: 'radius.azurecr.io/magpie:latest'
      }
    }
  }

  resource b 'Container' = {
    name: 'b'
    properties: {
      container: {
        image: 'radius.azurecr.io/magpie:latest'
      }
    }
  }
}