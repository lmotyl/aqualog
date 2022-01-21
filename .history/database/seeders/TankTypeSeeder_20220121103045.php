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
            TankType::WIDTH => 400,
            TankType::HEIGHT => 250,
            TankType::DEPTH => 250,
            TankType::GLASS_THICKNESS => 3
        ]);
        TankType::create(
        [
            TankType::WIDTH => 800,
            TankType::HEIGHT => 400,
            TankType::DEPTH => 350,
            TankType::GLASS_THICKNESS => 6
        ]);
        TankType::create(
        [
            TankType::WIDTH => 1500,
            TankType::HEIGHT => 600,
            TankType::DEPTH => 500,
            TankType::GLASS_THICKNESS => 11
        ]
        );
    }
}
