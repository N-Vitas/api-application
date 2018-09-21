import React from "react";
import { Grid } from "material-ui";
import { AddAlert } from "@material-ui/icons";
import { RegularCard, CustomInput, ItemGrid, Button, Snackbar } from "components";
import { Logout, LoggeIn, SetAuth, GetAuth } from "../../variables/session";

class Login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      alert: {
        show: false,
        type: 'danger',
        message: 'Не верный логин или пароль',
      },
      username: GetAuth().username,
      password: '',
    }
    this.handleUser = this.handleUser.bind(this);
    this.handlePass = this.handlePass.bind(this);
  }
  login() {
    const { alert, username, password } = this.state;
    LoggeIn({ username, password }, () => this.setState({alert:{...alert, show: true}}));
  }
  handleUser(e) { this.setState({ username: e.target.value }); }
  handlePass(e) { this.setState({ password: e.target.value }); }
  render() {
    // SetAuth({ token: 'JHkhLJKHKJSDlkjhHLSJD5sSDF64' });
    const { alert, username, password } = this.state;
    return (
      <Grid container direction="row" justify="center" alignItems="center">
        <ItemGrid xs={4} sm={12} md={4}>
          <RegularCard
            cardTitle={"Авторизация"+alert.show}
            cardSubtitle="Для начала работы пожалуйста авторизуйтесь"
            content={
              <div>
                <Grid container>
                  <ItemGrid xs={12} sm={12} md={12}>
                    <CustomInput
                      labelText="Логин"
                      id="username"
                      formControlProps={{
                        fullWidth: true,
                        onChange:this.handleUser,
                      }}
                    />
                  </ItemGrid>
                  <ItemGrid xs={12} sm={12} md={12}>
                    <CustomInput
                      labelText="Пароль"
                      id="password"
                      type="password"
                      formControlProps={{
                        fullWidth: true,
                        onChange:this.handlePass,
                      }}
                    />
                  </ItemGrid>
                  <ItemGrid xs={12} sm={12} md={12}>
                    <Button
                      fullWidth
                      color="primary"
                      onClick={() => this.login()}
                    >
                      Войти
                      </Button>
                    <Snackbar
                      place="tc"
                      color={alert.type}
                      icon={AddAlert}
                      message={alert.message}
                      open={alert.show}
                      closeNotification={() => this.setState({alert:{...alert, show: false}})}
                      close
                    />
                  </ItemGrid>
                </Grid>
              </div>
            }
          />
        </ItemGrid>
      </Grid>
    );
  }
}

export default Login;
