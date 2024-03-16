// @/pages/index.js
import React from 'react'
import Layout from '@/components/Layout/page';

export default function Home() {
  return (<>
    <Layout
      pageTitle=""
    >
      <div className="min-h-screen flex flex-col">
        <div className="m-auto">
          <h1 className="text-4xl"><b>Curso: </b>Sistemas Operativos 1</h1>
          <h2 className="text-4xl"><b>Nombre: </b>Douglas Alexander Soch Catalán</h2>
          <h2 className="text-4xl"><b>Carnet: </b>201807032</h2>
        </div>
      </div>
    </Layout>
  </>);
}