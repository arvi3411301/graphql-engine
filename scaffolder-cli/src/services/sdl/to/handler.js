const { getTypesSdl, getActionDefinitionSdl } = require('../../../shared/utils/sdlUtils');
const { deriveMutation } = require('../../derive/derive')

const handlePayload = (payload) => {

  const { actions, types, derive } = payload;
  const {
    mutation: toDeriveMutation,
    introspection_schema: introspectionSchema
  } = derive || {};

  const {
    name: deriveMutationName,
    action_name: actionName
  } = toDeriveMutation || {};

  const response = {
    body: null,
    status: 200
  };

  let actionSdl = '';
  let typesSdl = '';
  let actionSdlError, typesSdlError, deriveMutationError;

  if (actions) {
    try {
      actions.forEach(a => {
        actionSdl += getActionDefinitionSdl(a.name, a.arguments, a.output_type) + '\n';
      })
    } catch (e) {
      actionSdlError = e;
    }
  }

  if (types) {
    try {
      typesSdl = getTypesSdl(types);
    } catch (e) {
      typesSdlError = e;
    }
  }

  let sdl = `${actionSdl}\n\n${typesSdl}`;

  if (deriveMutationName) {
    try {
      const derivation = deriveMutation(deriveMutationName, introspectionSchema, actionName);
      const derivedActionSdl = getActionDefinitionSdl(derivation.action.name, derivation.action.arguments, derivation.action.output_type);
      const derivedTypesSdl = getTypesSdl(derivation.types);
      sdl = `${derivedActionSdl}\n\n${derivedTypesSdl}\n\n${sdl}`
    } catch (e) {
      deriveMutationError = e;
    }
  }

  if (actionSdlError) {
    response.body = {
      error: 'invalid actions definition'
    };
    response.status = 400;
    return response;
  }

  if (deriveMutationError) {
    response.body = {
      error: `could not derive mutation: ${deriveMutationError.message}`
    };
    response.status = 400;
    return response;
  }

  if (typesSdlError) {
    response.body = {
      error: 'invalid types'
    };
    response.status = 400;
    return response;
  }

  response.body = {
    sdl: {
      complete: sdl
    }
  };

  return response;

}

const requestHandler = (payload) => {

  const {
    body, status
  } = handlePayload(payload)

  return JSON.stringify(body);

}

module.exports = requestHandler;
module.exports.handlePayload = handlePayload;
