import { css, cx } from '@emotion/css'
import axios from 'axios'
import { useRouter } from 'next/router'
import React, { useState } from 'react'
import AppButton from '../templates/components/AppButton'
import AuthForm from '../templates/components/AuthForm'
import TextInput from '../templates/components/TextInput'
import PageTemplate from '../templates/PageTemplate'

const oauth2 = () => {
    const router = useRouter()

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [token, setToken] = useState("")

    const [consentData, setConsentData] = useState(undefined)

    const seoHeadData = {
        title: "GoAwth | Oauth2.0 Portal",
        description: "PoC OpenIDConnect & Oauth2.0 Authorization Server using Golang Echo, GORM, JWT, & MySQL"
    }
    
    return (
        <PageTemplate seoHead={seoHeadData}>
            <div className={oauth2Styles.body}>
                <div className={oauth2Styles.edge} />
                <div className={oauth2Styles.loginBody}>
                    <AuthForm subtitle={"Oauth2.0 Login Portal"}>
                        {consentData === undefined && (
                            <>
                                <TextInput
                                    password={false}
                                    placeholder="Username"
                                    style={{width: 225}}
                                    value={username}
                                    onChange={(e) => setUsername(e.target.value)}
                                />
                                <TextInput
                                    password={true}
                                    placeholder="Password"
                                    style={{width: 225}}
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                />
                                <AppButton 
                                    text={"Login"} 
                                    style={{marginTop: 20, width: 215}} 
                                    onClick={() => {
                                        loginApi({username, password}, (data) => {
                                            setToken(data.accessToken)
                                            setPassword("")
                                            oauthGetConsentApi({
                                                applicationId: router.query.applicationid,
                                                scope: router.query.scopes,
                                                grantType: router.query.granttype,
                                                playerUsername: username,
                                                token: data.accessToken,
                                            }, (data) => {
                                                setConsentData(data)
                                            })
                                        })
                                    }}
                                />  
                            </>
                        )}
                        {consentData !== undefined && (
                            <>
                                <div className={oauth2Styles.oauthImages}>
                                    <img
                                        src={consentData.developerApplicationImageUrl}
                                        alt=""
                                        style={{width: 30, height: 30, borderRadius: 80}}
                                    />
                                    <div className={oauth2Styles.oauthImagesDivider}/>
                                    <img
                                        src={consentData.playerImageUrl}
                                        alt=""
                                        style={{width: 30, height: 30, borderRadius: 80}}
                                    />
                                </div>
                                <div className={oauth2Styles.oauthScopeDefCtr}>
                                    <label className={cx(oauth2Styles.oauthScopeDef, oauth2Styles.oauthScopeDefTitle)}>Permissions Requested:</label>
                                    {Object.keys(consentData.scopeDefinitions).map((s, key) => {
                                        return (
                                            <label key={key} className={oauth2Styles.oauthScopeDef}>{consentData.scopeDefinitions[s]}</label>
                                        )
                                    })}
                                </div>
                                <AppButton
                                    text={"Agree"} 
                                    style={{marginTop: 20, width: 200}} 
                                />
                                <AppButton
                                    text={"Cancel"} 
                                    style={{marginTop: 20, width: 200, backgroundColor: "#E0144C"}} 
                                />
                            </>
                        )}
                    </AuthForm>
                </div>
                <div className={oauth2Styles.edge} />
            </div>
        </PageTemplate>
    )
}

export default oauth2

const oauth2Styles = {
    body: css`
        height: calc(100vh - 72px);
        display: flex;
    `,
    edge: css`
        flex: 1;
    `,
    loginBody: css`
        padding: 24px;
        display: flex;
        flex-direction: column;
        width: 400px;
    `,
    oauthImages: css`
        display: flex;
        justify-content: space-between;
        width: 150px;
        align-items: center;
        margin: 24px 0;
    `,
    oauthImagesDivider: css`
        border-top: 1px solid rgba(255, 255, 255, 0.05);
        width: 50px;
    `,
    oauthScopeDef: css`
        color: white;
        margin-top: 6px;
        font-size: 12px;
        font-weight: 200;
    `,
    oauthScopeDefCtr: css`
        background-color: rgb(15,15,15);
        padding: 18px;
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 200px;
        border-radius: 12px;
    `,
    oauthScopeDefTitle: css`
        font-weight: 400;
        color: #4CAF50;
        margin-bottom: 12px;
    `
}

const loginApi = ({username, password}, callback) => {
    axios.post("http://localhost:1323/player/login", {
        username,
        password
    }).then((response) => {
        console.log(response.data)
        callback(response.data.data)
    }).catch((response) => {
        console.log(response)
    })
}

const oauthGetConsentApi = ({applicationId, scope, grantType, playerUsername, token}, callback) => {
    const reqBody = {
        developerApplicationId: applicationId,
        scope,
        grantType,
        playerUsername
    }
    const reqHeader = {
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }

    console.log(reqBody, reqHeader)
    
    axios.post("http://localhost:1323/oauth2/get-consent", reqBody, reqHeader)
    .then((response) => {
        console.log(response.data)
        callback(response.data.data)
    }).catch((response) => {
        console.log(response)
    })
}