import { css } from '@emotion/css'
import React, { useState } from 'react'
import AppButton from '../templates/components/AppButton'
import AuthForm from '../templates/components/AuthForm'
import TextInput from '../templates/components/TextInput'
import PageTemplate from '../templates/PageTemplate'

const player = () => {
    const [login, setLogin] = useState(false)

    const seoHeadData = {
        title: "GoAwth | Player Portal",
        description: "PoC OpenIDConnect & Oauth2.0 Authorization Server using Golang Echo, GORM, JWT, & MySQL"
    }
    const authFormSubtitle = (login) => login ? "Login to Player Portal." : "Signup to Player Portal."

    const LoginModeLink = ({login}) => {
        if (login) {
            return (
                <label onClick={() => setLogin(false)} className={playerStyles.loginAlt}>Don't have an account? <p className={playerStyles.loginAltBold}>Signup</p></label>
            )
        }
        return (
            <label onClick={() => setLogin(true)} className={playerStyles.loginAlt}>Already have an account? <p className={playerStyles.loginAltBold}>Login</p></label>
        )
    }

    return (
        <PageTemplate seoHead={seoHeadData}>
            <div className={playerStyles.body}>
                <div className={playerStyles.documentationBody}></div>
                <div className={playerStyles.loginBody}>
                    <AuthForm subtitle={authFormSubtitle(login)}>
                        {login ? (
                            <>
                                <TextInput
                                    password={false}
                                    placeholder="Fullname"
                                    style={{width: 225}}
                                />
                                <TextInput
                                    password={true}
                                    placeholder="Password"
                                    style={{width: 225}}
                                />
                                <AppButton text={"Login"} style={{marginTop: 20, width: 215}} />
                            </>
                        ): (
                            <>
                                <TextInput
                                    password={false}
                                    placeholder="Fullname"
                                    style={{width: 225}}
                                />
                                <TextInput
                                    password={false}
                                    placeholder="Username"
                                    style={{width: 225}}
                                />
                                <TextInput
                                    password={false}
                                    placeholder="ImageUrl"
                                    style={{width: 225}}
                                />
                                <TextInput
                                    password={true}
                                    placeholder="Password"
                                    style={{width: 225}}
                                />
                                <AppButton text={"Signup"} style={{marginTop: 20, width: 215}} />
                            </>
                        )}
                        <LoginModeLink login={login}/>
                    </AuthForm>
                </div>
            </div>
        </PageTemplate>
    )
}

export default player

const playerStyles = {
    body: css`
        height: calc(100vh - 72px);
        display: flex;
    `,
    documentationBody: css`
        display: flex;
        flex-direction: column;
        padding: 24px;
        flex: 1;
        padding: 24px;
    `,
    loginBody: css`
        padding: 24px;
        display: flex;
        flex-direction: column;
        width: 400px;
    `,
    loginAlt: css`
        margin-top: 20px;
        font-size: 14px;
        font-weight: 200;
        display: flex;

        :hover {
            font-style: italic;
            cursor: pointer;
        }
    `,
    loginAltBold: css`
        font-weight: 600;
        color: #4CAF50;
        margin-left: 4px;
    `
}