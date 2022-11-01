import { css } from '@emotion/css'
import axios from 'axios'
import React, { useEffect, useState } from 'react'
import AppButton from '../templates/components/AppButton'
import AuthForm from '../templates/components/AuthForm'
import TextInput from '../templates/components/TextInput'
import PageTemplate from '../templates/PageTemplate'

function Developer() {
    const [login, setLogin] = useState(false)

    const [developerName, setDeveloperName] = useState("")
    const [password, setPassword] = useState("")

    const [developerToken, setDeveloperToken] = useState("")

    const resetState = () => {
        setDeveloperName("")
        setPassword("")
        setDeveloperToken("")
    }

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

    const authOnClick = (isLoginMode) => () => {
        if (isLoginMode) {
            devLoginApi({developerName, password}, (data) => {
                setPassword("");
                setDeveloperToken(data.accessToken)
            })
        } else {
            devSignupApi({developerName, password}, () => {
                resetState();
                setLogin(true);
            })
        }
    }

    return (
        <PageTemplate seoHead={seoHeadData}>
            {developerToken === "" ? (
                <div className={developerStyles.body}>
                    <div className={developerStyles.documentationBody}>
                        <DeveloperDocs/>
                    </div>
                    <div className={developerStyles.loginBody}>
                        <AuthForm subtitle={authFormSubtitle(login)}>
                            <TextInput
                                password={false}
                                placeholder="DeveloperName"
                                style={{width: 225}}
                                value={developerName}
                                onChange={(e) => setDeveloperName(e.target.value)}
                            />
                            <TextInput
                                password={true}
                                placeholder="Password"
                                style={{width: 225}}
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                            <AppButton 
                                text={login ? "Login" : "Signup"} 
                                style={{marginTop: 20, width: 215}} 
                                onClick={authOnClick(login)}
                            />
                            <LoginModeLink login={login}/>
                        </AuthForm>
                    </div>
                </div>
            ) : (
                <DeveloperPortal developerName={developerName} developerToken={developerToken} />
            )}
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
                    <li style={{marginBottom: 12}}>Create a developer account</li>
                    <li style={{marginBottom: 12}}>Create an application: remember to input the success & failed redirects</li>
                    <li style={{marginBottom: 12}}>In your application, redirect your login button to <code>{"http://localhost:3000/oauth2?applicationId=<appId>&scopes=<scopes>&grantType=<grantType>"}</code></li>
                    <li style={{marginBottom: 12}}>Hit a post request from your server side to this endpoint <code>{"http://localhost:1323/oauth2/authenticate-grant"}</code> to obtain your access token</li>
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

const devSignupApi = ({developerName, password}, callback) => {
    axios.post("http://localhost:1323/developer/sign-up", {
        developerName,
        password
    }).then((response) => {
        console.log(response.data)
        callback(response.data.data)
    }).catch((response) => {
        console.log(response)
    })
}

const devLoginApi = ({developerName, password}, callback) => {
    axios.post("http://localhost:1323/developer/login", {
        developerName,
        password
    }).then((response) => {
        console.log(response.data)
        callback(response.data.data)
    }).catch((response) => {
        console.log(response)
    })
}

const DeveloperPortal = ({developerToken, developerName}) => {

    const [devApps, setDevApps] = useState(undefined)
    const [activeIndex, setActiveIndex] = useState(0)

    useEffect(() => {
        devAppsListApi({developerName, developerToken}, (data) => setDevApps(data))
    }, [developerName, developerToken])

    if (devApps === undefined) {
        return <></>
    }

    return (
        <div className={developerPortalStyles.body}>
            <div className={developerPortalStyles.contentBody}>
                <div className={developerDocsStyles.body}>
                    <div className={developerDocsStyles.navBody}>
                        {devApps.applications.map(s => s.name).map((name, index) => (
                            <label 
                                key={index} 
                                className={developerDocsStyles.navButton}
                                onClick={() => setActiveIndex(index)}
                            >{name}</label>
                        ))}
                    </div>
                    <div className={developerDocsStyles.docsBody}>
                        <h4 className={developerDocsStyles.docsHeader}>{devApps.applications[activeIndex].name}</h4>
                        <div style={{height: 24}} />
                        <div className={developerPortalStyles.contentKeyValueBody}>
                            <div className={developerPortalStyles.contentKeyBody}>ApplicationId</div>
                            <div className={developerPortalStyles.contentValueBody}>{devApps.applications[activeIndex].developerApplicationId}</div>
                        </div>
                        <div className={developerPortalStyles.contentKeyValueBody}>
                            <div className={developerPortalStyles.contentKeyBody}>ApplicationName</div>
                            <div className={developerPortalStyles.contentValueBody}>{devApps.applications[activeIndex].name}</div>
                        </div>
                        <div className={developerPortalStyles.contentKeyValueBody}>
                            <div className={developerPortalStyles.contentKeyBody}>Secret</div>
                            <div className={developerPortalStyles.contentValueBody}>{devApps.applications[activeIndex].secret}</div>
                        </div>
                        <div className={developerPortalStyles.contentKeyValueBody}>
                            <div className={developerPortalStyles.contentKeyBody}>LogoUrl</div>
                            <div className={developerPortalStyles.contentValueBody}>{devApps.applications[activeIndex].logoUrl}</div>
                        </div>
                        <div className={developerPortalStyles.contentKeyValueBody}>
                            <div className={developerPortalStyles.contentKeyBody}>SuccessRedirectUri</div>
                            <div className={developerPortalStyles.contentValueBody}>{devApps.applications[activeIndex].successRedirectUri}</div>
                        </div>
                        <div className={developerPortalStyles.contentKeyValueBody}>
                            <div className={developerPortalStyles.contentKeyBody}>FailedRedirectUri</div>
                            <div className={developerPortalStyles.contentValueBody}>{devApps.applications[activeIndex].failedRedirectUri}</div>
                        </div>
                    </div>
                </div>
            </div>
            <div className={developerPortalStyles.formBody}>
                <p>create new app form</p>
            </div>
        </div>
    )
}

const developerPortalStyles = {
    body: css`
        height: calc(100vh - 72px);
        display: flex;
    `,
    contentBody: css`
        display: flex;
        flex-direction: column;
        padding: 24px;
        padding-right: 0;
        flex: 1;
    `,
    contentKeyValueBody: css`
        display: flex;
        align-items: center;
    `,
    contentKeyBody: css`
        padding: 12px;
        display: flex;
        align-items: center;
        background-color: rgb(10,10,10);
        border: 1px solid rgb(25,25,25);
        flex: 0.3;
    `,
    contentValueBody: css`
        padding: 12px;
        display: flex;
        align-items: center;
        justify-content: right;
        background-color: rgb(15,15,15);
        flex: 0.7;
    `,
    formBody: css`
        padding: 24px;
        display: flex;
        flex-direction: column;
        width: 400px;
    `,
}

const devAppsListApi = ({developerName, developerToken}, callback) => {
    axios.post("http://localhost:1323/dev-apps/list", {
        developerName,
        page: 1,
        pageSize: 10
    }, {
        headers: {
            "Authorization": `Bearer ${developerToken}`
        }
    }).then((response) => {
        console.log(response.data)
        callback(response.data.data)
    }).catch((response) => {
        console.log(response)
    })
}