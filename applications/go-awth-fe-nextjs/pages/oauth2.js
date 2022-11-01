import { css } from '@emotion/css'
import React from 'react'
import AppButton from '../templates/components/AppButton'
import AuthForm from '../templates/components/AuthForm'
import TextInput from '../templates/components/TextInput'
import PageTemplate from '../templates/PageTemplate'

const oauth2 = () => {
    const seoHeadData = {
        title: "GoAwth | Oauth2.0 Portal",
        description: "PoC OpenIDConnect & Oauth2.0 Authorization Server using Golang Echo, GORM, JWT, & MySQL"
    }
    
    return (
        <PageTemplate seoHead={seoHeadData}>
            <div className={oauth2Styles.body}>
                <div className={oauth2Styles.edge} />
                <div className={oauth2Styles.loginBody}>
                    <AuthForm subtitle={"Oauth2.0 Consent Portal"}>
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
    `
}