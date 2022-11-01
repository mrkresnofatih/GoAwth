import { css } from '@emotion/css'
import React from 'react'
import PageTemplate from '../templates/PageTemplate'

function Developer() {
    const seoHeadData = {
        title: "GoAwth | Developer Portal",
        description: "PoC OpenIDConnect & Oauth2.0 Authorization Server using Golang Echo, GORM, JWT, & MySQL"
    }

    return (
        <PageTemplate seoHead={seoHeadData}>
            <div className={developerStyles.body}>
                <div className={developerStyles.documentationBody}>hello</div>
                <div className={developerStyles.loginBody}>world</div>
            </div>
        </PageTemplate>
    )
}

export default Developer

const developerStyles = {
    body: css`
        height: calc(100vh - 72px);
        display: flex;
        background-color: blueviolet;
    `,
    documentationBody: css`
        display: flex;
        flex-direction: column;
        padding: 24px;
        flex: 1;
        padding: 24px;
        background-color: skyblue;
    `,
    loginBody: css`
        padding: 24px;
        display: flex;
        flex-direction: column;
        width: 400px;
        background-color: pink;
    `
}