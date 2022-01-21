<?php

namespace Database\Seeders;

use App\Models\TankType;
use Illuminate\Database\Seeder;

class TankTypeSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        TankType::create(            [
            'width' => 400,
            'height' => 250,
            'depth' => 250,
            'glass_thickness' => 3
        ],
        [
            'width' => 800,
            'height' => 400,
            'depth' => 350,
            'glass_thickness' => 6
        ],
        [
            'width' => 1500,
            'height' => 600,
            'depth' => 500,
            'glass_thickness' => 11
        ]
        );
    }
}
