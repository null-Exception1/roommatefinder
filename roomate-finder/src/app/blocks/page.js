"use client"
import Navbar from "@/components/navbar";
import Card from "@/components/card";
import Breadcumbs from "@/components/breadcrumbs";
import { useEffect, useState } from "react";

export default function Blocks() {
  const baseURL = process.env.NEXT_PUBLIC_API_URL;

  const url = `${baseURL}/blocks`;
  console.log(url)
  const [loaded, setLoaded] = useState(false)
  const [loaded_data, setLoadedData] = useState([])

  useEffect(() => {
    fetch(url).then(response => response.json()).then(data => { setLoaded(true); setLoadedData(data); console.log("Data:", loaded, data); })
  }, [])
  return (
    <div className="flex justify-center">
      <Navbar />

      <div className="flex flex-col m-20 flex-wrap w-11/12 border-2 border-black rounded-3xl">
        <div className="pt-10 pl-10">
          <fieldset className="fieldset max-w-xs text-black">
            <legend className="fieldset-legend text-black">Campus</legend>

            <select
              defaultValue="Main Campus"
              className="select bg-white text-black border-gray-300 focus:border-blue-500"
            >
              <option disabled={true}>Pick a campus</option>
              <option className="bg-white text-black">Main Campus</option>
              <option className="bg-white text-black">Bangalore</option>
              <option className="bg-white text-black">Jaipur</option>
            </select>
          </fieldset>
        </div>

        <Breadcumbs crumbs={["Home", "Blocks"]} links={["/", "/blocks"]} />

        <h1 className="pl-10 pt-2 text-2xl">Blocks</h1>

        <div className="grid grid-cols-3 p-5 *:gap-16 gap-16 ml-10">
          {loaded ? (
            Object.entries(loaded_data).map(([blockID, value]) => (
              <Card
                key={blockID}
                BlockID={blockID}
                PartialCount={value.Partial}
                FullCount={value.Full}
              />
            ))

          ) : (<div></div>)}
        </div>

      </div>

    </div>);
}