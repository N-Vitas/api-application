import axios from 'axios';

const state = {
  token: '',
  login: '',
  role: '',
  fullname: '',
  date: '',
}
const API = 'http://127.0.0.1:8080/'; 
const get = () => { return JSON.parse(window.localStorage.getItem('session') || state ); }
const set = (auth) => { window.localStorage.setItem('session', JSON.stringify(auth)); }
export const GenerateUrl = (uri) => {
  switch (uri) {
    case 'login':
      return `${API}api/token`;
    default:
      return `${API}${uri}`;
  } 
}
export const GetAuth = () => get() || {} ;
export const SetAuth = (auth) => set({...state, ...auth});
export const GetToken = () => get().token;
export const IsAuth = () => get().token != '';
export const Logout = () => set(state);
export const LoggeIn = (auth, cb) => {
  console.log('LoggeIn', auth, cb);
  axios.post(GenerateUrl('login'), auth).then(res => {
    console.log(res.data);
    cb();
  }).catch((response) => {
    console.log({ Message: response.message });
    cb();
  });
}