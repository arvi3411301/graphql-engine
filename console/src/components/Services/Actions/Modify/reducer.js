import defaultState from './state';

const SET_MODIFY_STATE = 'Actions/Modify/SET_MODIFY_STATE';
export const setModifyState = state => ({
  type: SET_MODIFY_STATE,
  state,
});

const SET_ACTION_HANDLER = 'Actions/Modify/SET_ACTION_HANDLER';
export const setActionHandler = handler => ({
  type: SET_ACTION_HANDLER,
  handler,
});

const SET_ACTION_KIND = 'Actions/Modify/SET_ACTION_KIND';
export const setActionKind = kind => ({
  type: SET_ACTION_KIND,
  kind,
});

const SET_ACTION_DEFINITION = 'Actions/Add/SET_ACTION_DEFINITION';
export const setActionDefinition = (sdl, error, timer, ast) => ({
  type: SET_ACTION_DEFINITION,
  definition: { sdl, error, timer, ast },
});

const SET_TYPE_DEFINITION = 'Actions/Add/SET_TYPE_DEFINITION';
export const setTypeDefinition = (sdl, error = null, timer, ast) => ({
  type: SET_TYPE_DEFINITION,
  definition: { sdl, error, timer, ast },
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
    case SET_ACTION_HANDLER:
      return {
        ...state,
        handler: action.handler,
      };
    case SET_ACTION_KIND:
      return {
        ...state,
        kind: action.kind,
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
    case SET_ACTION_DEFINITION:
      return {
        ...state,
        actionDefinition: {
          ...action.definition,
          sdl:
            action.definition.sdl !== null
              ? action.definition.sdl
              : state.actionDefinition.sdl,
        },
      };
    case SET_TYPE_DEFINITION:
      return {
        ...state,
        typeDefinition: {
          ...action.definition,
          sdl:
            action.definition.sdl !== null
              ? action.definition.sdl
              : state.typeDefinition.sdl,
        },
      };
    default:
      return state;
  }
};

export default reducer;