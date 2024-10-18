import React from 'react';

const Services = () => {
    return (
        <div className="pt-28 px-4 md:px-6 md:pt-20 pb-6">
            <div className="container mx-auto">
                <h1 className="text-3xl font-semibold text-zinc-600 mb-6">Our Services</h1>
                <p className="text-gray-600 mb-8">We offer a wide range of Water Quality services to meet your needs.</p>

                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    <div className="p-6 border rounded-lg shadow-md">
                        <h2 className="text-xl font-semibold text-zinc-600">Primary Purpose</h2>
                        <p className="text-gray-600 mt-2">Our primary purpose is to provide comprehensive and preventive measures and services for your healthy throuhg clean water.</p>
                    </div>

                    <div className="p-6 border rounded-lg shadow-md">
                        <h2 className="text-xl font-semibold text-zinc-600">Specialty</h2>
                        <p className="text-gray-600 mt-2">We have a team of specialists to address specific water issues, from water and enevironment department.</p>
                    </div>

                    <div className="p-6 border rounded-lg shadow-md">
                        <h2 className="text-xl font-semibold text-zinc-600">Time</h2>
                        <p className="text-gray-600 mt-2">Our emergency department is open 24/7 to provide immediate water supply services.</p>
                    </div>

                    <div className="p-6 border rounded-lg shadow-md">
                        <h2 className="text-xl font-semibold text-zinc-600">Quality Water</h2>
                        <p className="text-gray-600 mt-2">We prioritize quality and offer  services to support better well-being.</p>
                    </div>

                    {/* Add more services as needed */}
                </div>
            </div>
        </div>
    );
};

export default Services;
