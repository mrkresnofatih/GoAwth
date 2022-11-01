import { css } from '@emotion/css'
import React, { useState } from 'react'
import AppButton from '../templates/components/AppButton'
import AuthForm from '../templates/components/AuthForm'
import TextInput from '../templates/components/TextInput'
import PageTemplate from '../templates/PageTemplate'

function Developer() {
    const [login, setLogin] = useState(false)

    const seoHeadData = {
        title: "GoAwth | Developer Portal",
        description: "PoC OpenIDConnect & Oauth2.0 Authorization Server using Golang Echo, GORM, JWT, & MySQL"
    }
    const authFormSubtitle = (login) => login ? "Login to Developer Portal." : "Signup to Developer Portal."

    const LoginModeLink = ({login}) => {
        if (login) {
            return (
                <label onClick={() => setLogin(false)} className={developerStyles.loginAlt}>Don't have a developer account? <p className={developerStyles.loginAltBold}>Signup</p></label>
            )
        }
        return (
            <label onClick={() => setLogin(true)} className={developerStyles.loginAlt}>Have a developer account? <p className={developerStyles.loginAltBold}>Login</p></label>
        )
    }

    return (
        <PageTemplate seoHead={seoHeadData}>
            <div className={developerStyles.body}>
                <div className={developerStyles.documentationBody}>
                    <DeveloperDocs/>
                </div>
                <div className={developerStyles.loginBody}>
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

export default Developer

const developerStyles = {
    body: css`
        height: calc(100vh - 72px);
        display: flex;
    `,
    documentationBody: css`
        display: flex;
        flex-direction: column;
        padding: 24px;
        padding-right: 0;
        flex: 1;
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

const DeveloperDocs = () => {
    return (
        <div className={developerDocsStyles.body}>
            <div className={developerDocsStyles.navBody}>
                <label className={developerDocsStyles.navButton}>Getting Started</label>
            </div>
            <div className={developerDocsStyles.docsBody}>
                <h4 className={developerDocsStyles.docsHeader}>Getting Started</h4>
                <ol className={developerDocsStyles.docsList}>
                    <li>Create a developer account</li>
                    <li>Create an application: remember to input the success & failed redirects</li>
                    <li>In your application, redirect your login button to <code>{"http://localhost:3000/oauth2?applicationId=<appId>&scopes=<scopes>&grantType=<grantType>"}</code></li>
                </ol>
                
            </div>
        </div>
    )
}

const developerDocsStyles = {
    body: css`
        height: 100%;
        background-color: rgb(25,25,25);
        border-top: 10px solid #4CAF50;
        display: flex;
    `,
    navBody: css`
        display: flex;
        flex-direction: column;
        flex: 0.2;
        overflow-y: auto;
        -ms-overflow-style: none; /* for Internet Explorer, Edge */
        scrollbar-width: none; /* for Firefox */
        overflow-y: scroll; 

        ::-webkit-scrollbar {
            display: none; /* for Chrome, Safari, and Opera */
        }
    `,
    navButton: css`
        padding: 18px;
        display: flex;
        align-items: center;
        border-bottom: 1px solid rgba(255, 255, 255, 0.1);

        :hover {
            cursor: pointer;
            background-color: rgba(255, 255, 255, 0.05)
        }
    `,
    docsBody: css`
        padding: 24px;
        display: flex;
        flex-direction: column;
        flex: 0.8;
        border-left: 1px solid rgba(255, 255, 255, 0.1);
    `,
    docsHeader: css`
        font-size: 36px;
        font-weight: 600;
    `,
    docsList: css`
        margin-top: 24px;
        font-weight: 200;
    `
}