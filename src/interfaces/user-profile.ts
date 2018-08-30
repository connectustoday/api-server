import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IUserProfile {
	@prop({ required: true })
	schema_version: string;

	@prop()
	interests?: Array<string>;

	@prop()
	biography?: string;

	@prop()
	education?: string; //TODO

	@prop()
	quote?: string;

	@prop()
	current_residence?: string;

	@prop()
	certifications?: string; //TODO
}
