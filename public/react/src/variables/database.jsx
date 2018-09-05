const State = {
  token: '',
  login: '',
  role: '',
  name: '',
}
const jsonGet = () => JSON.parse(window.localStorage.getItem('api')) || State;

const jsonSet = (data = State) => window.localStorage.setItem('api', JSON.stringify(data));

export const getToken = () => jsonGet().token;
export const isAuth = () => jsonGet().token != '';

const UserDb = {
  getToken,
  isAuth,
}

export default UserDb;