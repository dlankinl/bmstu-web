import axios from 'axios';

const API_URL = 'http://localhost:8081/api/v1/';
// const API_URL = process.env.API_URL;

class AuthService {
  login(user) {
    return axios
      .post(API_URL + 'login', {
        login: user.username,
        password: user.password
      })
      .then(response => {
        if (response.data.token) {
          axios.defaults.headers.common['Authorization'] = 'Bearer ' + response.data.token;
          localStorage.setItem('user', JSON.stringify(response.data.token));
        }

        return response.data;
      });
  }

  logout() {
    localStorage.removeItem('user');
  }

  register(user) {
    return axios.post(API_URL + 'signup', {
      login: user.username,
      password: user.password
    });
  }
}

export default new AuthService();
