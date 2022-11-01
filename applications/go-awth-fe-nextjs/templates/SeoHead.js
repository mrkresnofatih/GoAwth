import Head from 'next/head'
import React from 'react'

const SeoHead = ({title, description}) => {
  return (
    <Head>
        <title>{title}</title>
        <meta name="description" content={description} />
        <link rel="icon" href="/goawthlogo.png" />
    </Head>
  )
}

export default SeoHead