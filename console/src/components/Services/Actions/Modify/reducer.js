import defaultState from './state';

const SET_MODIFY_STATE = 'Actions/Modify/SET_MODIFY_STATE';
export const setModifyState = state => ({
  type: SET_MODIFY_STATE,
  state,
});

const SET_ACTION_NAME = 'Actions/Modify/SET_ACTION_NAME';
export const setActionName = name => ({
  type: SET_ACTION_NAME,
  name,
});

const SET_ACTION_WEBHOOK = 'Actions/Modify/SET_ACTION_WEBHOOK';
export const setActionWebhook = webhook => ({
  type: SET_ACTION_WEBHOOK,
  webhook,
});

const SET_ACTION_KIND = 'Actions/Modify/SET_ACTION_KIND';
export const setActionKind = kind => ({
  type: SET_ACTION_KIND,
  kind,
});

const SET_ACTION_ARGUMENTS = 'Actions/Modify/SET_ACTION_ARGUMENTS';
export const setActionArguments = args => ({
  type: SET_ACTION_ARGUMENTS,
  args,
});

const SET_ACTION_OUTPUT_TYPE = 'Actions/Modify/SET_ACTION_OUTPUT_TYPE';
export const setActionOutputType = outputType => ({
  type: SET_ACTION_OUTPUT_TYPE,
  outputType,
});

const SET_TYPES = 'Actions/Modify/SET_TYPES';
export const setTypes = types => ({
  type: SET_TYPES,
  types,
});

// used to set types, args and output types together
const SET_TYPES_BULK = 'Actions/Modify/SET_TYPES_BULK';
export const setTypesBulk = (types, args, outputType) => ({
  type: SET_TYPES_BULK,
  args,
  types,
  outputType,
});

const SET_FETCHING = 'Actions/Modify/SET_FETCHING';
export const setFetching = () => ({ type: SET_FETCHING });
const UNSET_FETCHING = 'Actions/Modify/UNSET_FETCHING';
export const unsetFetching = () => ({ type: UNSET_FETCHING });

const reducer = (state = defaultState, action) => {
  switch (action.type) {
    case SET_MODIFY_STATE:
      return {
        ...action.state,
      };
    case SET_ACTION_NAME:
      return {
        ...state,
        name: action.name,
      };
    case SET_ACTION_WEBHOOK:
      return {
        ...state,
        webhook: action.webhook,
      };
    case SET_ACTION_KIND:
      return {
        ...state,
        kind: action.kind,
      };
    case SET_ACTION_ARGUMENTS:
      return {
        ...state,
        arguments: action.args,
      };
    case SET_ACTION_OUTPUT_TYPE:
      return {
        ...state,
        outputType: action.outputType,
      };
    case SET_TYPES:
      return {
        ...state,
        types: action.types,
      };
    case SET_TYPES_BULK:
      return {
        ...state,
        types: action.types,
        arguments: action.args,
        outputType: action.outputType,
      };
    case SET_FETCHING:
      return {
        ...state,
        isFetching: true,
      };
    case UNSET_FETCHING:
      return {
        ...state,
        isFetching: false,
      };
    default:
      return state;
  }
};

export default reducer;
