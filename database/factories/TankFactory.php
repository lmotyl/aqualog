<?php

namespace Database\Factories;

use App\Models\Tank;
use App\Models\TankType;
use App\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;

class TankFactory extends Factory
{


    protected $model = Tank::class;

    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            Tank::NAME => $this->faker->word
        ];
    }
}
