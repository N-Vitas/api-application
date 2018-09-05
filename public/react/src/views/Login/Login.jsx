import React from "react";
import { Grid, InputLabel } from "material-ui";

import { RegularCard, CustomInput, ItemGrid } from "components";

function Login({ ...props }) {
  return (
    <Grid container direction="row" justify="center" alignItems="center">
      <ItemGrid xs={4} sm={12} md={4}>
        <RegularCard
          cardTitle="Авторизация"
          cardSubtitle="Для начала работы пожалуйста авторизуйтесь"
          content={
            <div>
            <Grid container>
                  <ItemGrid xs={12} sm={12} md={12}>
                    <CustomInput
                      labelText="Логин"
                      id="username"
                      formControlProps={{
                        fullWidth: true
                      }}
                    />
                  </ItemGrid>
                  <ItemGrid xs={12} sm={12} md={12}>
                    <CustomInput
                      labelText="Пароль"
                      id="password"
                      type="password"
                      formControlProps={{
                        fullWidth: true
                      }}
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

export default Login;
